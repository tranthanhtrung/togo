// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add("login", (email, password) => { ... })
//
//
// -- This is a child command --
// Cypress.Commands.add("drag", { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add("dismiss", { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite("visit", (originalFn, url, options) => { ... })

Cypress.Commands.add('login', function(baseUrl, user_id, password) {
  return cy.request({
    method: 'GET',
    url: baseUrl + `/login?user_id=${user_id}&password=${password}`
  }).then(resp => {
    expect(resp.status).to.eq(200)
    expect(resp.body.data).to.be.a('string')
  })
})

Cypress.Commands.add('deleteTask', function(baseUrl, token, id) {
  return cy.request({
    method: 'DELETE',
    url: baseUrl + '/tasks',
    headers: {
      'Authorization': token
    },
    body: {
      id: id
    }
  }).then(resp => {
    expect(resp.status).to.eq(200)
    expect(resp.body.id).eq(id)
  })
})

Cypress.Commands.add('addTask', function(baseUrl, token) {
  return cy.request({
    method: 'POST',
    url: baseUrl + '/tasks',
    headers: {
      'Authorization': token
    },
    body: {
      content: 'content'
    }
  })
})