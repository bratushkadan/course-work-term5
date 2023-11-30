import { Product, GetProductsPayload } from './types';

export const transformProduct: (product: Product) => Product = (product) => ({
  ...product,
  price: product.price / 100,
});

export const transformProductsSearchParams = (params: URLSearchParams) => {
  const payload = {
    sort: omitFalsy({
      by: params.get('order_by'),
      order: params.get('order_by'),
    }),
    filter: omitFalsy({
      like_name: params.get('filter.like_name'),
      store_id: params.get('filter.store_id'),
      store_name: params.get('filter.store_name'),
      category_id: params.get('filter.category_id'),
      min_height: params.get('filter.min_height'),
      max_height: params.get('filter.max_height'),
      min_price: params.get('filter.min_price'),
      max_price: params.get('filter.max_price'),
    }),
  } as GetProductsPayload;
  return payload;
};

export function omitFalsy(obj: Record<string, unknown>) {
  return Object.fromEntries(Object.entries(obj).filter(([_, val]) => Boolean(val)));
}
