# wikible
This is "wikible" project for creating wiki tree as code.

## Background
Every time when we start the new project repeatedly, we should create the wiki(confluence) page and it should have similar contents such as project information, architecture and so on.

Confluence supports the template for the page of contents but the entire structure is not supported.

Wikible is a project for reducing the repeated tasks and providing manage code for the entire structure of wiki for the project.

- You can define your wiki strucuture with count of dash("-")

wiki.template
```
MyProject 
- Onboarding Guide 
- MyProject
-- Notice 
--- Release Notes 
--- Organization & Members 
--- Terms 
--- Planning and Schedule 
---- Roadmap 
---- Schedule 
```
Based on this template, you can create the wiki tree of your own project.




## Authentication
- Wikible supports `Basic access authentication` and you can use your id/password when you excute the command "apply".
In the case of the Confluence, You should use `API Token` instead of password.
- https://confluence.atlassian.com/cloud/api-tokens-938839638.html


## Features
### Template
Wikible supports a template and you can see the example of template in `/template/ops.template`.
You can define your own template such as `development.template`, `projectmanage.template` and so on.
It follows the `golang template syntax(https://golang.org/pkg/text/template/)`


### Project
You can define your projects using template. you can see the examples in `/template/project/bixby.yaml`
It contains the below information
```
1. name: project name
2. template: template location
3. vars: variables which defined in the template
4. version: project template version
5. writer: writer of this document
```

### Commands
1. Plan
- for checking the numbering of wiki tree and binding variables with templates
- options
  - *-p* : project template file
    - sample project template file: template/project/bixby.yaml
    - It contains the template file location and binding variables
  - *-i* : root page id 
- examples
  - `./wikible plan -p template/project/bixby.yaml`
  - `./wikible plan -p template/project/bixby.yaml -a https://rpkim.atlassian.net/wiki -i 208584739`
  - `./wikible plan -p template/project/bixby.yaml -a https://rpkim.atlassian.net/wiki`
  
2. Apply
- creating the wiki tree with the template
- options
  - *-p* : project template file
    - sample project template file: template/project/bixby.yaml
    - It contains the template file location and binding variables
  - *-a* : wiki api address
    - ex> `https://mobilderndhub.sec.samsung.net/wiki`
  - *-i* : root page id 
- examples
  - `./wikible apply -p template/project/test.yaml`
  - `./wikible apply -p template/project/test.yaml -a https://rpkim.atlassian.net/wiki`
  - `./wikible apply -p template/project/test.yaml -a https://rpkim.atlassian.net/wiki -i 208584739`

## TBD
- get "curl http://a/wiki/rest/api/content/487427029/child/page | jq '.results[] | {title:.title, id:.id}'
- https://a/wiki/rest/api/content/821303369?expand=body.storage
