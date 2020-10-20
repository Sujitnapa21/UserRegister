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
    EntSexEdges,
    EntSexEdgesFromJSON,
    EntSexEdgesFromJSONTyped,
    EntSexEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSex
 */
export interface EntSex {
    /**
     * 
     * @type {EntSexEdges}
     * @memberof EntSex
     */
    edges?: EntSexEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntSex
     */
    id?: number;
    /**
     * Sexname holds the value of the "sexname" field.
     * @type {string}
     * @memberof EntSex
     */
    sexname?: string;
}

export function EntSexFromJSON(json: any): EntSex {
    return EntSexFromJSONTyped(json, false);
}

export function EntSexFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSex {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntSexEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'sexname': !exists(json, 'sexname') ? undefined : json['sexname'],
    };
}

export function EntSexToJSON(value?: EntSex | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntSexEdgesToJSON(value.edges),
        'id': value.id,
        'sexname': value.sexname,
    };
}


