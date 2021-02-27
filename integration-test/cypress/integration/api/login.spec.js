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
})
