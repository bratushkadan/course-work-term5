import {useEffect, useState} from 'react'
import {useAuth, useCart, useFavorite} from '../stores'
import {api} from '../api'
import {alertError} from '../util/error'

export const RequestAppDataComponent: React.FC = () => {
  const token = useAuth(state => state.token)
  const [initDone, setInitDone] = useState(false)

  const setCart = useCart(state => state.setCart)
  const setFavorites = useFavorite(state => state.setFavorites)

  useEffect(() => {
    if (token === null) {
      return
    }
    if (initDone) {
      return
    }
    setInitDone(true)

    api.getCart(token).then(setCart).catch(alertError)
    api.getFavorites(token).then(setFavorites).catch(alertError)
  }, [token, setInitDone])

  return null
}