import {
  BASEURL,
  USER_DEFAULT
} from '../../common/commonAPI.js'


describe('[Login] Create Task', function() {
  let token

  before(() => {
    cy.login(BASEURL, USER_DEFAULT.user_id, USER_DEFAULT.password).then(response => {
      token = response.body.data
    })
  })

  it('Add a task for user', function() {
    cy.request({
      method: 'POST',
      url: BASEURL + '/tasks',
      headers: {
        'Authorization': token
      },
      timeout: 60000,
      body: {
        content: 'Some content'
      }
    }).then((resp) => {
      expect(resp.status).to.eq(200)
      expect(resp.body.data).to.have.all.keys('id', 'content', 'user_id', 'created_date')
      expect(resp.body.data.content).eq('Some content')

      cy.deleteTask(BASEURL, token, resp.body.data.id)
    })
  })
  
  it('Add a task quota exceed limit tasks', async function() {
    let promises = []
    for(let i = 0; i < 5; i++) {
      promises.push(cy.addTask(BASEURL, token))
    }

    Promise.all(promises).then(values => {
      cy.request({
        method: 'POST',
        url: BASEURL + '/tasks',
        headers: {
          'Authorization': token
        },
        timeout: 60000,
        body: {
          content: 'Some content'
        },
        failOnStatusCode: false,
      }).then((resp) => {
        expect(resp.status).to.eq(400)
        expect(resp.body.error).eq('Limit 5 task per day')
      })
      
      values.forEach(value => {
        expect(value.status).to.eq(200)
        expect(value.body.data).to.have.all.keys('id', 'content', 'user_id', 'created_date')
        cy.deleteTask(BASEURL, token, value.body.data.id)
      })
    })
  })

  it('Create task of a user in a day without token', function() {
    cy.request({
      method: 'POST',
      url: BASEURL + '/tasks',
      timeout: 60000,
      body: {
        content: 'Some content'
      },
      failOnStatusCode: false,
    }).then((resp) => {
      expect(resp.status).to.eq(401)
      expect(resp.body.error).eq('Unauthorized')
    })
  })

  it('Create task of a user in a day with invalid token', function() {
    cy.request({
      method: 'POST',
      url: BASEURL + '/tasks',
      headers: {
        'Authorization': 'test'
      },
      body: {
        content: 'Some content'
      },
      timeout: 60000,
      failOnStatusCode: false,
    }).then((resp) => {
      expect(resp.status).to.eq(401)
      expect(resp.body.error).eq('Unauthorized')
    })
  })

  it('Create task of a user without body content', function() {
    cy.request({
      method: 'POST',
      url: BASEURL + '/tasks',
      headers: {
        'Authorization': token
      },
      timeout: 60000,
      failOnStatusCode: false,
    }).then((resp) => {
      expect(resp.status).to.eq(500)
    })
  })

  it('Create task of a user with wrong url', function() {
    cy.request({
      method: 'POST',
      url: BASEURL + '/task',
      headers: {
        'Authorization': token
      },
      body: {
        content: 'Some content'
      },
      timeout: 60000,
      failOnStatusCode: false,
    }).then((resp) => {
      expect(resp.status).to.eq(404)
      expect(resp.body.error).eq('Not fould')
    })
  })
})

describe('[Login] List Tasks', function() {
  let token

  before(() => {
    cy.login(BASEURL, USER_DEFAULT.user_id, USER_DEFAULT.password).then(response => {
      token = response.body.data
    })
  })

  it('Get list tasks of a user in a day', function() {
    cy.request({
      method: 'GET',
      url: BASEURL + '/tasks?created_date=2021-02-27',
      headers: {
        'Authorization': token
      },
      timeout: 60000,
    }).then((resp) => {
      expect(resp.status).to.eq(200)
      resp.body.data.forEach(el => {
        expect(el).to.have.all.keys('id', 'content', 'user_id', 'created_date')
        expect(el.created_date).eq('2021-02-27')
      })
    })
  })

  it('Get list tasks of a user in a day without token', function() {
    cy.request({
      method: 'GET',
      url: BASEURL + '/tasks?created_date=2021-02-27',
      timeout: 60000,
      failOnStatusCode: false,
    }).then((resp) => {
      expect(resp.status).to.eq(401)
      expect(resp.body.error).eq('Unauthorized')
    })
  })

  it('Get list tasks of a user in a day with invalid token', function() {
    cy.request({
      method: 'GET',
      url: BASEURL + '/tasks?created_date=2021-02-27',
      headers: {
        'Authorization': 'test'
      },
      timeout: 60000,
      failOnStatusCode: false,
    }).then((resp) => {
      expect(resp.status).to.eq(401)
      expect(resp.body.error).eq('Unauthorized')
    })
  })

  it('Get list tasks of a user with wrong format date', function() {
    cy.request({
      method: 'GET',
      url: BASEURL + '/tasks?created_date=27-02-2021',
      headers: {
        'Authorization': token
      },
      timeout: 60000,
      failOnStatusCode: false,
    }).then((resp) => {
      expect(resp.status).to.eq(200)
      expect(resp.body.data).eq(null)
    })
  })

  it('Get list tasks of a user without param created_date', function() {
    cy.request({
      method: 'GET',
      url: BASEURL + '/tasks',
      headers: {
        'Authorization': token
      },
      timeout: 60000,
      failOnStatusCode: false,
    }).then((resp) => {
      expect(resp.status).to.eq(200)
      expect(resp.body.data).eq(null)
    })
  })

  it('Get list tasks of a user with wrong url', function() {
    cy.request({
      method: 'GET',
      url: BASEURL + '/task?created_date=2021-02-27',
      headers: {
        'Authorization': token
      },
      timeout: 60000,
      failOnStatusCode: false,
    }).then((resp) => {
      expect(resp.status).to.eq(404)
      expect(resp.body.error).eq('Not fould')
    })
  })
})