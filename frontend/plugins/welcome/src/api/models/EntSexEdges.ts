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

import { exists, mapValues } from '../runtime';
import {
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSexEdges
 */
export interface EntSexEdges {
    /**
     * User holds the value of the user edge.
     * @type {Array<EntUser>}
     * @memberof EntSexEdges
     */
    user?: Array<EntUser>;
}

export function EntSexEdgesFromJSON(json: any): EntSexEdges {
    return EntSexEdgesFromJSONTyped(json, false);
}

export function EntSexEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSexEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'user': !exists(json, 'user') ? undefined : ((json['user'] as Array<any>).map(EntUserFromJSON)),
    };
}

export function EntSexEdgesToJSON(value?: EntSexEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'user': value.user === undefined ? undefined : ((value.user as Array<any>).map(EntUserToJSON)),
    };
}


