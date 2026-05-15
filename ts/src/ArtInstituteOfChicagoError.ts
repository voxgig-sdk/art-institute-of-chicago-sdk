
import { Context } from './Context'


class ArtInstituteOfChicagoError extends Error {

  isArtInstituteOfChicagoError = true

  sdk = 'ArtInstituteOfChicago'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  ArtInstituteOfChicagoError
}

