import { Product, GetProductsPayload, OrderInfo } from './types';

export const transformProduct: (product: Product) => Product = (product) => ({
  ...product,
  price: product.price / 100,
});

export const transformProductsSearchParams = (params: URLSearchParams) => {
  const payload = {
    sort: omitFalsy({
      by: params.get('sort_by'),
      order: params.get('sort_order'),
    }),
    filter: omitFalsy({
      like_name: params.get('product_name'),
      store_id: params.get('store_id'),
      store_name: params.get('store_name'),
      category_id: params.get('category_id'),
      min_height: params.get('min_height'),
      max_height: params.get('max_height'),
      min_price: Math.trunc(Number(params.get('min_price')) * 100),
      max_price: Math.trunc(Number(params.get('max_price')) * 100),
    }),
  } as GetProductsPayload;
  return payload;
};

export function omitFalsy(obj: Record<string, unknown>) {
  return Object.fromEntries(Object.entries(obj).filter(([_, val]) => Boolean(val)));
}

export const localizeOrderStatus = (status: OrderInfo['status']): string => {
  switch (status) {
    case 'created':
      return 'Создан'
    case 'canceled':
      return 'Отменен'
    case 'completed':
      return 'Выполнен'
    case 'delivery':
      return 'Доставляется'
    case 'in_progress':
      return 'В сборке'
    case 'processed':
      return 'Собран и ожидает отправки'
    default:
    throw new Error('такого статуса заказа нет!')
  }
}
