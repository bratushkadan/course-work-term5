import {HTTPError} from 'ky'

import { Error } from '../api/types'

export async function alertError(err: unknown) {
  if (err instanceof HTTPError) {
    const errorJson: Error = await err.response.json()
    alert(`Ошибка: ${errorJson.error}`)
  }
}