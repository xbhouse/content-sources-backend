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
 * @interface ApiUpload
 */
export interface ApiUpload {
    /**
     * HREF to the unfinished upload, use with internal API
     * @type {string}
     * @memberof ApiUpload
     */
    href?: string;
    /**
     * SHA256 sum of the uploaded file
     * @type {string}
     * @memberof ApiUpload
     */
    sha256?: string;
    /**
     * Upload UUID, use with public API
     * @type {string}
     * @memberof ApiUpload
     */
    uuid?: string;
}

/**
 * Check if a given object implements the ApiUpload interface.
 */
export function instanceOfApiUpload(value: object): value is ApiUpload {
    return true;
}

export function ApiUploadFromJSON(json: any): ApiUpload {
    return ApiUploadFromJSONTyped(json, false);
}

export function ApiUploadFromJSONTyped(json: any, ignoreDiscriminator: boolean): ApiUpload {
    if (json == null) {
        return json;
    }
    return {
        
        'href': json['href'] == null ? undefined : json['href'],
        'sha256': json['sha256'] == null ? undefined : json['sha256'],
        'uuid': json['uuid'] == null ? undefined : json['uuid'],
    };
}

export function ApiUploadToJSON(json: any): ApiUpload {
    return ApiUploadToJSONTyped(json, false);
}

export function ApiUploadToJSONTyped(value?: ApiUpload | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'href': value['href'],
        'sha256': value['sha256'],
        'uuid': value['uuid'],
    };
}

