package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Wiki Struct
type Wiki struct {
	url       *url.URL
	basicAuth string
	client    *http.Client
}

//CreateWikiAPI is for the Wiki Object
func CreateWikiAPI(address string, basicAuth string) (*Wiki, error) {
	u, err := url.ParseRequestURI(address)
	if err != nil {
		return nil, err
	}

	if !strings.HasSuffix(u.Path, "/") {
		u.Path += "/"
	}

	u.Path += "rest/api"
	wiki := new(Wiki)
	wiki.url = u
	wiki.basicAuth = basicAuth
	wiki.client = &http.Client{}

	return wiki, nil
}

func (w *Wiki) sendRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("Accept", "application/json, */*")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", w.basicAuth)

	resp, err := w.client.Do(req)
	if err != nil {
		if resp != nil {
			res, err := ioutil.ReadAll(resp.Body)
			return res, err
		}
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusPartialContent:
		res, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return res, err
		}
		return res, nil
	case http.StatusNoContent, http.StatusResetContent:
		return nil, nil
	case http.StatusUnauthorized:
		return nil, fmt.Errorf("authentication failed")
	case http.StatusServiceUnavailable:
		return nil, fmt.Errorf("service is not available (%s)", resp.Status)
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("internal server error: %s", resp.Status)
	}

	if resp != nil {
		res, _ := ioutil.ReadAll(resp.Body)

		if strings.Contains(string(res[:]),"Basic auth with password is not allowed on this instance") {
			return res, fmt.Errorf("Please use the API Token instead of Password, https://id.atlassian.com/manage/api-tokens")
		} else {
			return res, fmt.Errorf("unknown response status %s", resp.Status)
		}
	}
	return nil, err
}

func (w *Wiki) contentEndpoint(contentID string) (*url.URL, error) {
	return url.ParseRequestURI(w.url.String() + "/content/" + contentID)
}

//GetPageContent is for getting Space Key for creating wiki
func (w *Wiki) GetPageContent(rootpageID string) (*PageContent, error) {
	contentEndPoint, err := w.contentEndpoint(rootpageID)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", contentEndPoint.String(), nil)
	if err != nil {
		fmt.Println("NewRequest Creation Error")
		fmt.Println(err)
		return nil, err
	}

	res, err := w.sendRequest(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var pageContent PageContent
	err = json.Unmarshal(res, &pageContent)
	if err != nil {
		fmt.Println("Json Umarshal Error")
		fmt.Println(err)
		return nil, err
	}

	return &pageContent, nil
}

//GetContent is for getting the wiki content from contentID
func (w *Wiki) GetContent(contentID string, expand []string) (*Content, error) {
	contentEndPoint, err := w.contentEndpoint(contentID)
	if err != nil {
		return nil, err
	}
	data := url.Values{}
	data.Set("expand", strings.Join(expand, ","))
	contentEndPoint.RawQuery = data.Encode()

	req, err := http.NewRequest("GET", contentEndPoint.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := w.sendRequest(req)
	if err != nil {
		return nil, err
	}

	var content Content
	err = json.Unmarshal(res, &content)
	if err != nil {
		return nil, err
	}

	return &content, nil
}

func createbody(title string, ancestorID string, space Space, body string) (*bytes.Buffer, error) {
	ancestor := Ancestor{
		AncestorID: ancestorID,
	}

	ancestors := Ancestors{
		ancestor,
	}

	var content ContentBody
	content.Storage.Representation = "storage"
	content.Storage.Value = body

	page := &ContentPage{
		Title:     title,
		Type:      "page",
		Ancestors: ancestors,
		Space:     space,
	}

	page.ContentBody = content


	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(page)

	return buf, nil
}

//CreateContent api
func (w *Wiki) CreateContent(title string, ancestorID string, space Space, body string) ([]byte, error) {
	data, err := createbody(title, ancestorID, space, body)

	addr, err := url.ParseRequestURI(w.url.String() + "/content/")

	req, err := http.NewRequest("POST", addr.String(), data)

	res, err := w.sendRequest(req)

	if err != nil {
		return res, err
	}

	return res, nil
}