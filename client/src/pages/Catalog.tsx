import {useEffect} from 'react'
import {api} from '../api'
import {useCatalog} from '../stores/catalog'
import {useShallow} from 'zustand/react/shallow'
import {Product} from '../components/catalog/Product'
import styled from 'styled-components'
import {useSearchParams} from 'react-router-dom'
import {transformProductsSearchParams} from '../api/util'

const CatalogWrapper = styled.div`
`

export const Catalog: React.FC = () => {
  const {products, setProducts} = useCatalog(useShallow(state => ({
    products: state.products,
    setProducts: state.setProducts,
  })))

  const [searchParams, setSearchParams] = useSearchParams()

  useEffect(() => {
    api.getProducts(
      transformProductsSearchParams(searchParams)
    ).then(setProducts)
  }, [searchParams])

  return (<CatalogWrapper>
  <h1>Каталог</h1>
  {products.map(product => <Product {...product} key={product.id}/>)}
  </CatalogWrapper>)
}