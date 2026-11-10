/* eslint-disable */
/* tslint:disable */
// @ts-nocheck
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface ActualitecoActualiteInternalCoreDomainCustomError {
  code?: string;
  /** @example "invalid request" */
  message?: string;
}

export interface ActualitecoActualiteInternalCoreDomainInternalError {
  code?: "INTERNAL_ERROR";
  message?: string;
}

export interface InternalFeaturesAuthTransportHttpLoginDTO {
  /** @example "test.user@email.com" */
  email: string;
  /** @example "hashpassword" */
  password: string;
}

export interface InternalFeaturesAuthTransportHttpLoginDTOResponse {
  token?: string;
  user_id?: string;
  username?: string;
}

export interface InternalFeaturesAuthTransportHttpRegisterDTO {
  /** @example "m4rkek@gmail.com" */
  email?: string;
  /** @example "m4rkek" */
  username?: string;
  /** @example "supersecretpass" */
  password?: string;
  /** @example "admin" */
  role?: string;
}

export interface InternalFeaturesAuthTransportHttpRegisterDTOResponse {
  /** @example "3fa85f64-5717-4562-b3fc-2c963f66afa6" */
  id?: string;
  /** @example "m4rkek" */
  username?: string;
  /** @example "m4rkek@gmail.com" */
  email?: string;
  /** @example "user" */
  role?: string;
  /** @example "2026-03-25T12:00:41.267Z" */
  createdAt?: string;
}
