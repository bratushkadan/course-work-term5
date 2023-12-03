import {useEffect, useState} from 'react'
import type {Store as IStore} from '../api/types'
import {api} from '../api'
import {Store} from '../components/stores/Store'

export const StoresPage: React.FC = () => {
  const [stores, setStores] = useState<IStore[]>()
  
  useEffect(() => {
    api.getStores().then(setStores)
  }, [])

  return <>
    {stores?.map(props => <Store {...props} key={props.id} isWithLinkToStore={true}/>)}
  </>
}