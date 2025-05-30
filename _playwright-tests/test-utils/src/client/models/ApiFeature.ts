/* tslint:disable */
/* eslint-disable */
/**
 * ContentSourcesBackend
 * The API for the repositories of the content sources that you can use to create and manage repositories between third-party applications and the [Red Hat Hybrid Cloud Console](https://console.redhat.com). With these repositories, you can build and deploy images using Image Builder for Cloud, on-Premise, and Edge. You can handle tasks, search for required RPMs, fetch a GPGKey from the URL, and list the features within applications. 
 *
 * The version of the OpenAPI document: v1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ApiFeature
 */
export interface ApiFeature {
    /**
     * Whether the current user can access the feature
     * @type {boolean}
     * @memberof ApiFeature
     */
    accessible?: boolean;
    /**
     * Whether the feature is enabled on the running server
     * @type {boolean}
     * @memberof ApiFeature
     */
    enabled?: boolean;
}

/**
 * Check if a given object implements the ApiFeature interface.
 */
export function instanceOfApiFeature(value: object): value is ApiFeature {
    return true;
}

export function ApiFeatureFromJSON(json: any): ApiFeature {
    return ApiFeatureFromJSONTyped(json, false);
}

export function ApiFeatureFromJSONTyped(json: any, ignoreDiscriminator: boolean): ApiFeature {
    if (json == null) {
        return json;
    }
    return {
        
        'accessible': json['accessible'] == null ? undefined : json['accessible'],
        'enabled': json['enabled'] == null ? undefined : json['enabled'],
    };
}

export function ApiFeatureToJSON(json: any): ApiFeature {
    return ApiFeatureToJSONTyped(json, false);
}

export function ApiFeatureToJSONTyped(value?: ApiFeature | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'accessible': value['accessible'],
        'enabled': value['enabled'],
    };
}

