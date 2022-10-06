class CPeopleList {
  template = null;

  constructor(itemplate) {
    this.template = itemplate
  }
  render(data) {
    if (!this.template) {
      return '<p>Error template not loaded</p>';
    }
    return this.template(data)
  }
}

function getPeopleList(Handlebars) {
  return fetch("/site/templates/people-list.html")
  .then(resp => resp.text())
  .then(templ => {
    var t = Handlebars.compile(templ);
    new CPeopleList(template);
  });
};
