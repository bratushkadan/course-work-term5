export type Error = {
  error: string;
};

export type User = {
  email: string;
  first_name: string;
  id: number;
  last_name: string;
  phone_number: string;
};

export type CreateUserPayload = {
  first_name: string;
  last_name: string;
  email: string;
  password: string;
  phone_number: string;
};
export type CreateUserResponse = User;

export type Token = {
  token: string;
};

export type GetTokenPayload = {
  email: string;
  password: string;
};

export type ValidateTokenPayload = Token;
export type ValidateTokenResponse = {
  valid: boolean;
};

export type Product = {
  id: number;
  store_name: string;
  store_id: number;
  name: string;
  description: string;
  image_url: string;
  price: number;
  created: number;
  category: {
    id: number;
    name: string;
    description: string;
  };
  min_height: number;
  max_height: number;
};

export type GetProductsPayload = Partial<{
  sort: Partial<{
    by: 'id' | 'price' | 'store_id' | 'max_height';
    order: 'asc' | 'desc';
  }>;
  filter: Partial<{
    like_name: string;
    store_id: number;
    category_id: number;
    min_height: number;
    max_height: number;
    min_price: number;
    max_price: number;
  }>;
}>;

export type Store = {
  id: number;
  name: string;
  email: string;
  phone_number: string;
  created: number;
};

export type Category = {
  id: number;
  name: string;
  description: string;
};

export type CartPosition = {
  product_id: number;
  quantity: number;
  name: string;
  description: string;
  price: number;
  image_url: string;
  category_id: number;
  category_name: string;
  store_id: number;
  store_name: string;
};

export type ChangeCartPositionPayload = {
  product_id: number;
  quantity: number;
};

export type MinimalCartPosition = {
  product_id: number;
  quantity: number;
};

export type Order = {
  id: number;
  status: 'created';
  user_id: number;
  created: number;
  status_modified: number;
  positions: Array<{
    product_id: number;
    quantity: number;
    name: string;
    description: string;
    price: number;
    image_url: string;
    category_id: number;
    category_name: string;
    store_id: number;
    store_name: string;
  }>;
};

export type OrderInfo = {
  id: number;
  status: 'created' | 'in_progress' | 'processed' | 'delivery' | 'canceled' | 'completed';
  user_id: number;
  created: number;
  status_modified: number;
};

export type Favorite = {
  product_id: number;
  store_id: number;
  store_name: string;
  name: string;
  description: string;
  image_url: string;
  price: number;
  category_id: number;
  category_name: string;
  added_favorite: number;
};

export type AddFavoritePayload = {
  product_id: number;
};

export type DeleteFavoritePayload = {
  product_id: number;
};

export type AddDeleteFavoriteResponse = {
  /**
   * product_id
   */
  id: number;
};

export type Review = {
  id: number;
  user_id: number;
  product_id: number;
  product_name: string;
  /**
   * min - 1
   * max - 10
   */
  rating: number;
  review_text: string;
  created: number;
  modified: number;
  user_name: string;
};

export type GetReviewsPayload = {
  user_id?: number;
  product_id?: number;
  store_id?: number;
};

export type AddReviewPayload = {
  product_id: number;
  rating: number;
  review_text: string;
};
export type EditReviewPayload = {
  product_id: number;
  rating: number;
  review_text: string;
};
export type DeleteReviewPayload = {
  product_id: number;
};
export type DeleteFavoriteResponse = {
  /**
   * product_id
   */
  id: number;
};
