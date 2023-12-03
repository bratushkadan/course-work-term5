import { useEffect, useLayoutEffect, useState } from 'react';
import { api } from '../api';
import { useCatalog } from '../stores/catalog';
import { useShallow } from 'zustand/react/shallow';
import { ProductComponent } from '../components/product/Product';
import styled from 'styled-components';
import { useSearchParams } from 'react-router-dom';
import { transformProductsSearchParams } from '../api/util';
import ReactModal from 'react-modal';
import { BlockBtn, BlockInput } from '../components/generic';
import { useDebounce } from 'use-debounce';

const CatalogWrapper = styled.div``;

export const Catalog: React.FC = () => {
  const { products, setProducts } = useCatalog(
    useShallow((state) => ({
      products: state.products,
      setProducts: state.setProducts,
    }))
  );

  const [isShowModal, setIsShowModal] = useState(false);

  const [searchParams, setSearchParams] = useSearchParams();
  const [debouncedSearchParams] = useDebounce(searchParams, 200);

  const [productNameInpVal, setProductNameInpVal] = useState(searchParams.get('product_name') || '');
  const [minPrice, setMinPrice] = useState<number | undefined>(Number(searchParams.get('min_price')) || undefined);
  const [maxPrice, setMaxPrice] = useState<number | undefined>(Number(searchParams.get('max_price')) || undefined);
  const [minHeight, setMinHeight] = useState<number | undefined>(Number(searchParams.get('min_height')) || undefined);
  const [maxHeight, setMaxHeight] = useState<number | undefined>(Number(searchParams.get('max_height')) || undefined);
  const [sortBy, setSortBy] = useState(searchParams.get('sort_by') || '');
  const [sortOrder, setSortOrder] = useState(searchParams.get('sort_order') || '');

  useEffect(() => {
    api.getProducts(transformProductsSearchParams(debouncedSearchParams)).then(setProducts);
  }, [debouncedSearchParams]);

  const handleSetSearchProductByNameParams = (e: React.ChangeEvent<HTMLInputElement>) => {
    const text = e.currentTarget.value.trim();
    setProductNameInpVal(e.currentTarget.value);
    if (text === '') {
      setSearchParams((params) => {
        params.delete('product_name');
        return params;
      });
    } else {
      setSearchParams((params) => {
        params.set('product_name', text);
        return params;
      });
    }
  };

  const handleChangeMinPrice = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.currentTarget.value === '') {
      setMinPrice(undefined);
      setSearchParams((params) => {
        params.delete('min_price');
        return params;
      });
    } else {
      setMinPrice(Number(e.currentTarget.value));
      setSearchParams((params) => {
        params.set('min_price', e.currentTarget.value);
        return params;
      });
    }
  };
  const handleChangeMaxPrice = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.currentTarget.value === '') {
      setMaxPrice(undefined);
      setSearchParams((params) => {
        params.delete('max_price');
        return params;
      });
    } else {
      setMaxPrice(Number(e.currentTarget.value));
      setSearchParams((params) => {
        params.set('max_price', e.currentTarget.value);
        return params;
      });
    }
  };
  const handleChangeMinHeight = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.currentTarget.value === '') {
      setMinHeight(undefined);
      setSearchParams((params) => {
        params.delete('min_height');
        return params;
      });
    } else {
      setMinHeight(Number(e.currentTarget.value));
      setSearchParams((params) => {
        params.set('min_height', e.currentTarget.value);
        return params;
      });
    }
  };
  const handleChangeMaxHeight = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.currentTarget.value === '') {
      setMaxHeight(undefined);
      setSearchParams((params) => {
        params.delete('max_height');
        return params;
      });
    } else {
      setMaxHeight(Number(e.currentTarget.value));
      setSearchParams((params) => {
        params.set('max_height', e.currentTarget.value);
        return params;
      });
    }
  };

  const handleSelectSortBy = (e: React.ChangeEvent<HTMLSelectElement>) => {
    if (e.currentTarget.value === "") {
      setSearchParams(params => {
        params.delete('sort_by')
        return params
      })
    } else {
      setSearchParams(params => {
        params.set('sort_by', e.currentTarget.value)
        return params
      })
    }
    console.log(e.currentTarget.value)
    setSortBy(e.currentTarget.value)
  }
  const handleSelectSortOrder = (e: React.ChangeEvent<HTMLSelectElement>) => {
    if (e.currentTarget.value === "") {
      setSearchParams(params => {
        params.delete('sort_order')
        return params
      })
    } else {
      setSearchParams(params => {
        params.set('sort_order', e.currentTarget.value)
        return params
      })
    }
    setSortOrder(e.currentTarget.value)
  }

  return (
    <CatalogWrapper>
      <h1>Каталог</h1>
      {/* <BlockBtn onClick={() => setIsShowModal(true)}>Выбрать фильтры поиска</BlockBtn> */}
      <label htmlFor="name">Поиск: </label>
      <input
        type="text"
        placeholder="введите название товара"
        value={productNameInpVal}
        onChange={handleSetSearchProductByNameParams}
      />
      <hr/>
      <label htmlFor="min_price">Минимальная цена: </label>
        <input type="number" placeholder="1000" value={minPrice} onChange={handleChangeMinPrice} />
        <br />
        <label htmlFor="min_price">Максимальная цена: </label>
        <input type="number" placeholder="20000" value={maxPrice} onChange={handleChangeMaxPrice} />
        <br />
        <label htmlFor="min_price">Минимальная высота (см): </label>
        <input type="number" placeholder="10" value={minHeight} onChange={handleChangeMinHeight} />
        <br />
        <label htmlFor="min_price">Максимальная высота (см): </label>
        <input type="number" placeholder="300" value={maxHeight} onChange={handleChangeMaxHeight} />
        <hr />
        <label htmlFor="sort_by">Порядок отображения: </label>
        <select name="sort_by" id="sort_by" value={sortBy} onChange={handleSelectSortBy}>
          <option value="">
            Не выбрано
          </option>
          <option value="id">По дате добавления растения</option>
          <option value="price">По стоимости растения</option>
          <option value="store_id">По дате регистрации продавца</option>
          <option value="max_height">По высоте растения</option>
        </select>
        <br />
        <label htmlFor="sort_order">Порядок сортировки:</label>
        <select name="sort_order" id="sort_order" value={sortOrder} onChange={handleSelectSortOrder}>
          <option value="">
            Не выбрано
          </option>
          <option value="asc">По возрастанию</option>
          <option value="desc">По убыванию</option>
        </select>
      {/* <ReactModal
        isOpen={isShowModal}
        shouldFocusAfterRender={true}
        shouldCloseOnOverlayClick={true}
        shouldCloseOnEsc={true}
        ariaHideApp={false}
        onRequestClose={() => setIsShowModal(false)}
        contentLabel="modal"
      >
        <label htmlFor="min_price">Минимальная цена: </label>
        <input type="number" placeholder="1000" value={minPrice} onChange={handleChangeMinPrice} />
        <br />
        <label htmlFor="min_price">Максимальная цена: </label>
        <input type="number" placeholder="20000" value={maxPrice} onChange={handleChangeMaxPrice} />
        <br />
        <label htmlFor="min_price">Минимальная высота (см): </label>
        <input type="number" placeholder="10" value={minHeight} onChange={handleChangeMinHeight} />
        <br />
        <label htmlFor="min_price">Максимальная высота (см): </label>
        <input type="number" placeholder="300" value={maxHeight} onChange={handleChangeMaxHeight} />
        <hr />
        <label htmlFor="sort_by">Порядок отображения: </label>
        <select name="sort_by" id="sort_by" onSelect={(e) => console.log(e.currentTarget.value)}>
          <option value="" selected={true}>
            Не выбрано
          </option>
          <option value="id">Сначала новые</option>
          <option value="price">Сначала самые недорогие</option>
          <option value="store_id">Сначала от самых старых продавцов</option>
          <option value="max_height">Сначала самые высокие</option>
        </select>
        <br />
        <label htmlFor="sort_order">Порядок сортировки:</label>
        <select name="sort_order" id="sort_order">
          <option value="" selected={true}>
            Не выбрано
          </option>
          <option value="desc">По возрастанию</option>
          <option value="asc">По убыванию</option>
        </select>
        <br />
        <BlockBtn
          style={{
            position: 'absolute',
            top: '1rem',
            right: '1rem',
          }}
          onClick={() => setIsShowModal(false)}
        >
          Закрыть окно
        </BlockBtn>
      </ReactModal> */}
      {products.length === 0 ? (
        <h2>Нет совпадений товаров</h2>
      ) : (
        products.map((product) => (
          <ProductComponent
            {...product}
            key={product.id}
            isWithProductLink={true}
            isAddToCartControls={true}
            isAddToFavoriteControls={true}
          />
        ))
      )}
    </CatalogWrapper>
  );
};
