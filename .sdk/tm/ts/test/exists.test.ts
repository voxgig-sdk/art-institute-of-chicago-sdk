
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { ArtInstituteOfChicagoSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = ArtInstituteOfChicagoSDK.test()
    equal(testsdk instanceof ArtInstituteOfChicagoSDK, true,
      'ArtInstituteOfChicagoSDK.test() must return a client synchronously')
  })

})
