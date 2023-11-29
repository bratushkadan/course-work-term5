export type Error = {
  error: string
}

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
  phone_number: string
};
export type CreateUserResponse = User

export type Token = {
  token: string
}

export type GetTokenPayload = {
  email: string
  password: string
}

export type ValidateTokenPayload = Token
export type ValidateTokenResponse = {
  valid: boolean
}