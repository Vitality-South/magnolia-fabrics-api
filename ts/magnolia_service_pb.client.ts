/* eslint-disable */
// @generated by protobuf-ts 2.8.2 with parameter eslint_disable,add_pb_suffix,long_type_number
// @generated from protobuf file "magnolia_service.proto" (package "magnoliafabrics", syntax proto3)
// tslint:disable
//
// Copyright (c) 2022 Vitality South, LLC <chris@vitalitysouth.com>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { MagnoliaFabricsService } from "./magnolia_service_pb";
import type { GetCleaningCodesResponse } from "./magnolia_service_pb";
import type { GetCleaningCodesRequest } from "./magnolia_service_pb";
import type { GetAllFabricTaxonomyResponse } from "./magnolia_service_pb";
import type { GetAllFabricTaxonomyRequest } from "./magnolia_service_pb";
import type { GetFabricBySKUResponse } from "./magnolia_service_pb";
import type { GetFabricBySKURequest } from "./magnolia_service_pb";
import type { GetFabricByNameResponse } from "./magnolia_service_pb";
import type { GetFabricByNameRequest } from "./magnolia_service_pb";
import type { GetFabricByIDResponse } from "./magnolia_service_pb";
import type { GetFabricByIDRequest } from "./magnolia_service_pb";
import type { GetAllInventoryResponse } from "./magnolia_service_pb";
import type { GetAllInventoryRequest } from "./magnolia_service_pb";
import type { GetAllFabricsWithoutInventoryResponse } from "./magnolia_service_pb";
import type { GetAllFabricsWithoutInventoryRequest } from "./magnolia_service_pb";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { GetAllFabricsResponse } from "./magnolia_service_pb";
import type { GetAllFabricsRequest } from "./magnolia_service_pb";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service magnoliafabrics.MagnoliaFabricsService
 */
export interface IMagnoliaFabricsServiceClient {
    /**
     * @generated from protobuf rpc: GetAllFabrics(magnoliafabrics.GetAllFabricsRequest) returns (magnoliafabrics.GetAllFabricsResponse);
     */
    getAllFabrics(input: GetAllFabricsRequest, options?: RpcOptions): UnaryCall<GetAllFabricsRequest, GetAllFabricsResponse>;
    /**
     * @generated from protobuf rpc: GetAllFabricsWithoutInventory(magnoliafabrics.GetAllFabricsWithoutInventoryRequest) returns (magnoliafabrics.GetAllFabricsWithoutInventoryResponse);
     */
    getAllFabricsWithoutInventory(input: GetAllFabricsWithoutInventoryRequest, options?: RpcOptions): UnaryCall<GetAllFabricsWithoutInventoryRequest, GetAllFabricsWithoutInventoryResponse>;
    /**
     * @generated from protobuf rpc: GetAllInventory(magnoliafabrics.GetAllInventoryRequest) returns (magnoliafabrics.GetAllInventoryResponse);
     */
    getAllInventory(input: GetAllInventoryRequest, options?: RpcOptions): UnaryCall<GetAllInventoryRequest, GetAllInventoryResponse>;
    /**
     * @generated from protobuf rpc: GetFabricByID(magnoliafabrics.GetFabricByIDRequest) returns (magnoliafabrics.GetFabricByIDResponse);
     */
    getFabricByID(input: GetFabricByIDRequest, options?: RpcOptions): UnaryCall<GetFabricByIDRequest, GetFabricByIDResponse>;
    /**
     * @generated from protobuf rpc: GetFabricByName(magnoliafabrics.GetFabricByNameRequest) returns (magnoliafabrics.GetFabricByNameResponse);
     */
    getFabricByName(input: GetFabricByNameRequest, options?: RpcOptions): UnaryCall<GetFabricByNameRequest, GetFabricByNameResponse>;
    /**
     * @generated from protobuf rpc: GetFabricBySKU(magnoliafabrics.GetFabricBySKURequest) returns (magnoliafabrics.GetFabricBySKUResponse);
     */
    getFabricBySKU(input: GetFabricBySKURequest, options?: RpcOptions): UnaryCall<GetFabricBySKURequest, GetFabricBySKUResponse>;
    /**
     * @generated from protobuf rpc: GetAllFabricTaxonomy(magnoliafabrics.GetAllFabricTaxonomyRequest) returns (magnoliafabrics.GetAllFabricTaxonomyResponse);
     */
    getAllFabricTaxonomy(input: GetAllFabricTaxonomyRequest, options?: RpcOptions): UnaryCall<GetAllFabricTaxonomyRequest, GetAllFabricTaxonomyResponse>;
    /**
     * @generated from protobuf rpc: GetCleaningCodes(magnoliafabrics.GetCleaningCodesRequest) returns (magnoliafabrics.GetCleaningCodesResponse);
     */
    getCleaningCodes(input: GetCleaningCodesRequest, options?: RpcOptions): UnaryCall<GetCleaningCodesRequest, GetCleaningCodesResponse>;
}
/**
 * @generated from protobuf service magnoliafabrics.MagnoliaFabricsService
 */
export class MagnoliaFabricsServiceClient implements IMagnoliaFabricsServiceClient, ServiceInfo {
    typeName = MagnoliaFabricsService.typeName;
    methods = MagnoliaFabricsService.methods;
    options = MagnoliaFabricsService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: GetAllFabrics(magnoliafabrics.GetAllFabricsRequest) returns (magnoliafabrics.GetAllFabricsResponse);
     */
    getAllFabrics(input: GetAllFabricsRequest, options?: RpcOptions): UnaryCall<GetAllFabricsRequest, GetAllFabricsResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetAllFabricsRequest, GetAllFabricsResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetAllFabricsWithoutInventory(magnoliafabrics.GetAllFabricsWithoutInventoryRequest) returns (magnoliafabrics.GetAllFabricsWithoutInventoryResponse);
     */
    getAllFabricsWithoutInventory(input: GetAllFabricsWithoutInventoryRequest, options?: RpcOptions): UnaryCall<GetAllFabricsWithoutInventoryRequest, GetAllFabricsWithoutInventoryResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetAllFabricsWithoutInventoryRequest, GetAllFabricsWithoutInventoryResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetAllInventory(magnoliafabrics.GetAllInventoryRequest) returns (magnoliafabrics.GetAllInventoryResponse);
     */
    getAllInventory(input: GetAllInventoryRequest, options?: RpcOptions): UnaryCall<GetAllInventoryRequest, GetAllInventoryResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetAllInventoryRequest, GetAllInventoryResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetFabricByID(magnoliafabrics.GetFabricByIDRequest) returns (magnoliafabrics.GetFabricByIDResponse);
     */
    getFabricByID(input: GetFabricByIDRequest, options?: RpcOptions): UnaryCall<GetFabricByIDRequest, GetFabricByIDResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetFabricByIDRequest, GetFabricByIDResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetFabricByName(magnoliafabrics.GetFabricByNameRequest) returns (magnoliafabrics.GetFabricByNameResponse);
     */
    getFabricByName(input: GetFabricByNameRequest, options?: RpcOptions): UnaryCall<GetFabricByNameRequest, GetFabricByNameResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetFabricByNameRequest, GetFabricByNameResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetFabricBySKU(magnoliafabrics.GetFabricBySKURequest) returns (magnoliafabrics.GetFabricBySKUResponse);
     */
    getFabricBySKU(input: GetFabricBySKURequest, options?: RpcOptions): UnaryCall<GetFabricBySKURequest, GetFabricBySKUResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetFabricBySKURequest, GetFabricBySKUResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetAllFabricTaxonomy(magnoliafabrics.GetAllFabricTaxonomyRequest) returns (magnoliafabrics.GetAllFabricTaxonomyResponse);
     */
    getAllFabricTaxonomy(input: GetAllFabricTaxonomyRequest, options?: RpcOptions): UnaryCall<GetAllFabricTaxonomyRequest, GetAllFabricTaxonomyResponse> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetAllFabricTaxonomyRequest, GetAllFabricTaxonomyResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetCleaningCodes(magnoliafabrics.GetCleaningCodesRequest) returns (magnoliafabrics.GetCleaningCodesResponse);
     */
    getCleaningCodes(input: GetCleaningCodesRequest, options?: RpcOptions): UnaryCall<GetCleaningCodesRequest, GetCleaningCodesResponse> {
        const method = this.methods[7], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetCleaningCodesRequest, GetCleaningCodesResponse>("unary", this._transport, method, opt, input);
    }
}
