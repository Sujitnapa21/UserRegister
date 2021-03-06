/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    EntRole,
    EntRoleFromJSON,
    EntRoleToJSON,
    EntSex,
    EntSexFromJSON,
    EntSexToJSON,
    EntTitle,
    EntTitleFromJSON,
    EntTitleToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserToJSON,
} from '../models';

export interface CreateRoleRequest {
    role: EntRole;
}

export interface CreateSexRequest {
    sex: EntSex;
}

export interface CreateTitleRequest {
    title: EntTitle;
}

export interface CreateUserRequest {
    user: EntUser;
}

export interface DeleteUserRequest {
    id: number;
}

export interface GetRoleRequest {
    id: number;
}

export interface GetSexRequest {
    id: number;
}

export interface GetTitleRequest {
    id: number;
}

export interface GetUserRequest {
    id: number;
}

export interface ListRoleRequest {
    limit?: number;
    offset?: number;
}

export interface ListSexRequest {
    limit?: number;
    offset?: number;
}

export interface ListTitleRequest {
    limit?: number;
    offset?: number;
}

export interface ListUserRequest {
    limit?: number;
    offset?: number;
}

export interface UpdateUserRequest {
    id: number;
    user: EntUser;
}

/**
 * 
 */
export class DefaultApi extends runtime.BaseAPI {

    /**
     * Create role
     * Create role
     */
    async createRoleRaw(requestParameters: CreateRoleRequest): Promise<runtime.ApiResponse<EntRole>> {
        if (requestParameters.role === null || requestParameters.role === undefined) {
            throw new runtime.RequiredError('role','Required parameter requestParameters.role was null or undefined when calling createRole.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/roles`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntRoleToJSON(requestParameters.role),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntRoleFromJSON(jsonValue));
    }

    /**
     * Create role
     * Create role
     */
    async createRole(requestParameters: CreateRoleRequest): Promise<EntRole> {
        const response = await this.createRoleRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create sex
     * Create sex
     */
    async createSexRaw(requestParameters: CreateSexRequest): Promise<runtime.ApiResponse<EntSex>> {
        if (requestParameters.sex === null || requestParameters.sex === undefined) {
            throw new runtime.RequiredError('sex','Required parameter requestParameters.sex was null or undefined when calling createSex.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/sexs`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntSexToJSON(requestParameters.sex),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntSexFromJSON(jsonValue));
    }

    /**
     * Create sex
     * Create sex
     */
    async createSex(requestParameters: CreateSexRequest): Promise<EntSex> {
        const response = await this.createSexRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create title
     * Create title
     */
    async createTitleRaw(requestParameters: CreateTitleRequest): Promise<runtime.ApiResponse<EntTitle>> {
        if (requestParameters.title === null || requestParameters.title === undefined) {
            throw new runtime.RequiredError('title','Required parameter requestParameters.title was null or undefined when calling createTitle.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/titles`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntTitleToJSON(requestParameters.title),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntTitleFromJSON(jsonValue));
    }

    /**
     * Create title
     * Create title
     */
    async createTitle(requestParameters: CreateTitleRequest): Promise<EntTitle> {
        const response = await this.createTitleRaw(requestParameters);
        return await response.value();
    }

    /**
     * Create user
     * Create user
     */
    async createUserRaw(requestParameters: CreateUserRequest): Promise<runtime.ApiResponse<EntUser>> {
        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling createUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: EntUserToJSON(requestParameters.user),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntUserFromJSON(jsonValue));
    }

    /**
     * Create user
     * Create user
     */
    async createUser(requestParameters: CreateUserRequest): Promise<EntUser> {
        const response = await this.createUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * get user by ID
     * Delete a user entity by ID
     */
    async deleteUserRaw(requestParameters: DeleteUserRequest): Promise<runtime.ApiResponse<object>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling deleteUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse<any>(response);
    }

    /**
     * get user by ID
     * Delete a user entity by ID
     */
    async deleteUser(requestParameters: DeleteUserRequest): Promise<object> {
        const response = await this.deleteUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * get role by ID
     * Get a role entity by ID
     */
    async getRoleRaw(requestParameters: GetRoleRequest): Promise<runtime.ApiResponse<EntRole>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getRole.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/roles/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntRoleFromJSON(jsonValue));
    }

    /**
     * get role by ID
     * Get a role entity by ID
     */
    async getRole(requestParameters: GetRoleRequest): Promise<EntRole> {
        const response = await this.getRoleRaw(requestParameters);
        return await response.value();
    }

    /**
     * get sex by ID
     * Get a sex entity by ID
     */
    async getSexRaw(requestParameters: GetSexRequest): Promise<runtime.ApiResponse<EntSex>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getSex.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/sexs/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntSexFromJSON(jsonValue));
    }

    /**
     * get sex by ID
     * Get a sex entity by ID
     */
    async getSex(requestParameters: GetSexRequest): Promise<EntSex> {
        const response = await this.getSexRaw(requestParameters);
        return await response.value();
    }

    /**
     * get title by ID
     * Get a title entity by ID
     */
    async getTitleRaw(requestParameters: GetTitleRequest): Promise<runtime.ApiResponse<EntTitle>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getTitle.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/titles/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntTitleFromJSON(jsonValue));
    }

    /**
     * get title by ID
     * Get a title entity by ID
     */
    async getTitle(requestParameters: GetTitleRequest): Promise<EntTitle> {
        const response = await this.getTitleRaw(requestParameters);
        return await response.value();
    }

    /**
     * get user by ID
     * Get a user entity by ID
     */
    async getUserRaw(requestParameters: GetUserRequest): Promise<runtime.ApiResponse<EntUser>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling getUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntUserFromJSON(jsonValue));
    }

    /**
     * get user by ID
     * Get a user entity by ID
     */
    async getUser(requestParameters: GetUserRequest): Promise<EntUser> {
        const response = await this.getUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * list role entities
     * List role entities
     */
    async listRoleRaw(requestParameters: ListRoleRequest): Promise<runtime.ApiResponse<Array<EntRole>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/roles`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntRoleFromJSON));
    }

    /**
     * list role entities
     * List role entities
     */
    async listRole(requestParameters: ListRoleRequest): Promise<Array<EntRole>> {
        const response = await this.listRoleRaw(requestParameters);
        return await response.value();
    }

    /**
     * list sex entities
     * List sex entities
     */
    async listSexRaw(requestParameters: ListSexRequest): Promise<runtime.ApiResponse<Array<EntSex>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/sexs`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntSexFromJSON));
    }

    /**
     * list sex entities
     * List sex entities
     */
    async listSex(requestParameters: ListSexRequest): Promise<Array<EntSex>> {
        const response = await this.listSexRaw(requestParameters);
        return await response.value();
    }

    /**
     * list title entities
     * List title entities
     */
    async listTitleRaw(requestParameters: ListTitleRequest): Promise<runtime.ApiResponse<Array<EntTitle>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/titles`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntTitleFromJSON));
    }

    /**
     * list title entities
     * List title entities
     */
    async listTitle(requestParameters: ListTitleRequest): Promise<Array<EntTitle>> {
        const response = await this.listTitleRaw(requestParameters);
        return await response.value();
    }

    /**
     * list user entities
     * List user entities
     */
    async listUserRaw(requestParameters: ListUserRequest): Promise<runtime.ApiResponse<Array<EntUser>>> {
        const queryParameters: runtime.HTTPQuery = {};

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(EntUserFromJSON));
    }

    /**
     * list user entities
     * List user entities
     */
    async listUser(requestParameters: ListUserRequest): Promise<Array<EntUser>> {
        const response = await this.listUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * update user by ID
     * Update a user entity by ID
     */
    async updateUserRaw(requestParameters: UpdateUserRequest): Promise<runtime.ApiResponse<EntUser>> {
        if (requestParameters.id === null || requestParameters.id === undefined) {
            throw new runtime.RequiredError('id','Required parameter requestParameters.id was null or undefined when calling updateUser.');
        }

        if (requestParameters.user === null || requestParameters.user === undefined) {
            throw new runtime.RequiredError('user','Required parameter requestParameters.user was null or undefined when calling updateUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters.id))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: EntUserToJSON(requestParameters.user),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => EntUserFromJSON(jsonValue));
    }

    /**
     * update user by ID
     * Update a user entity by ID
     */
    async updateUser(requestParameters: UpdateUserRequest): Promise<EntUser> {
        const response = await this.updateUserRaw(requestParameters);
        return await response.value();
    }

}
