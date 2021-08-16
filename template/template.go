package template

import "fmt"

// Deployment define the standard deploy process of a project
type Deployment interface {
	Deploy()
}

type template interface {
	pre()
	download()
	deploy()
	post()
}

type defaultTemplate struct {
	template
	source string
	target string
}

func newTemplate(t template) *defaultTemplate {
	return &defaultTemplate{template: t}
}

func (t *defaultTemplate) pre() {
	fmt.Println("default pre")
}

func (t *defaultTemplate) download() {
	fmt.Println("default download")
}

func (t *defaultTemplate) deploy() {
	fmt.Println("default deploy")
}

func (t *defaultTemplate) post() {
	fmt.Println("default post")
}

func (t *defaultTemplate) Deploy() {
	t.template.pre()
	t.template.download()
	t.template.deploy()
	t.template.post()
}

type DeployProject struct {
	*defaultTemplate
}

func NewProject(source, target string) Deployment {
	p := &DeployProject{}
	t := newTemplate(p)
	t.source = source
	t.target = target
	p.defaultTemplate = t
	return p
}

func (p *DeployProject) download() {
	fmt.Println("download", p.source)
}

func (p *DeployProject) deploy() {
	fmt.Println("deploy", p.target)
}
