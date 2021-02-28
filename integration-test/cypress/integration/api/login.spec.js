import {
  BASEURL,
} from '../../common/commonAPI.js'

describe('[Login] Test Login', function() {
  it('Login with valid user_id, password', function() {
    cy.request({
      method: 'GET',
      url: BASEURL + '/login?user_id=firstUser&password=example',
      timeout: 60000,
    }).then((resp) => {
      expect(resp.status).to.eq(200)
      expect(resp.body.data).to.be.a('string')
    })
  })
  
  it('Login with invalid user_id', function() {
    cy.request({
      method: 'GET',
      url: BASEURL + '/login?user_id=firstUse&password=example',
      timeout: 60000,
      failOnStatusCode: false,
    }).then((resp) => {
      expect(resp.status).to.eq(401)
      expect(resp.body.error).eq('incorrect user_id/pwd')
    })
  })

  it('Login with invalid password', function() {
    cy.request({
      method: 'GET',
      url: BASEURL + '/login?user_id=firstUser&password=exampler',
      timeout: 60000,
      failOnStatusCode: false,
    }).then((resp) => {
      expect(resp.status).to.eq(401)
      expect(resp.body.error).eq('incorrect user_id/pwd')
    })
  })

  it('Login with missing params', function() {
    cy.request({
      method: 'GET',
      url: BASEURL + '/login',
      timeout: 60000,
      failOnStatusCode: false,
    }).then((resp) => {
      expect(resp.status).to.eq(401)
      expect(resp.body.error).eq('incorrect user_id/pwd')
    })
  })
})
