/* eslint-disable */
// @generated by protobuf-ts 2.8.1 with parameter eslint_disable,add_pb_suffix,long_type_number
// @generated from protobuf file "api_fabric.proto" (package "magnoliafabrics", syntax proto3)
// tslint:disable
//
// Copyright (c) 2022 Vitality South, LLC <chris@vitalitysouth.com>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Inventory } from "./api_inventory_pb";
/**
 * @generated from protobuf message magnoliafabrics.Fabric
 */
export interface Fabric {
    /**
     * @gotags: dynamo:"FabricSKU"
     *
     * @generated from protobuf field: string sku = 1;
     */
    sku: string;
    /**
     * @gotags: dynamo:"FabricProductCode"
     *
     * @generated from protobuf field: string product_code = 2;
     */
    productCode: string;
    /**
     * @gotags: dynamo:"FabricPattern"
     *
     * @generated from protobuf field: string pattern = 3;
     */
    pattern: string;
    /**
     * @gotags: dynamo:"FabricPrimaryColor"
     *
     * @generated from protobuf field: string color = 4;
     */
    color: string;
    /**
     * @gotags: dynamo:"FabricPatternColorCombo"
     *
     * @generated from protobuf field: string pattern_color_combo = 5;
     */
    patternColorCombo: string;
    /**
     * @gotags: dynamo:"FabricStatus"
     *
     * @generated from protobuf field: string status = 6;
     */
    status: string;
    /**
     * @gotags: dynamo:"FabricIntroDate"
     *
     * @generated from protobuf field: string intro_date = 7;
     */
    introDate: string;
    /**
     * @gotags: dynamo:"FabricContents"
     *
     * @generated from protobuf field: string contents = 8;
     */
    contents: string;
    /**
     * @gotags: dynamo:"FabricWidth"
     *
     * @generated from protobuf field: string width = 9;
     */
    width: string;
    /**
     * @gotags: dynamo:"FabricDirection"
     *
     * @generated from protobuf field: string direction = 10;
     */
    direction: string;
    /**
     * @gotags: dynamo:"FabricHRepeat"
     *
     * @generated from protobuf field: string h_repeat = 11;
     */
    hRepeat: string;
    /**
     * @gotags: dynamo:"FabricVRepeat"
     *
     * @generated from protobuf field: string v_repeat = 12;
     */
    vRepeat: string;
    /**
     * @gotags: dynamo:"FabricCleaningCode"
     *
     * @generated from protobuf field: string cleaning_code = 13;
     */
    cleaningCode: string;
    /**
     * @gotags: dynamo:"FabricDoubleRubs"
     *
     * @generated from protobuf field: string double_rubs = 14;
     */
    doubleRubs: string;
    /**
     * @gotags: dynamo:"FabricMisc"
     *
     * @generated from protobuf field: string misc = 15;
     */
    misc: string;
    /**
     * @gotags: dynamo:"FabricDisclaimer"
     *
     * @generated from protobuf field: string disclaimer = 16;
     */
    disclaimer: string;
    /**
     * @gotags: dynamo:"FabricUses"
     *
     * @generated from protobuf field: repeated string uses = 17;
     */
    uses: string[];
    /**
     * @gotags: dynamo:"FabricDesigns"
     *
     * @generated from protobuf field: repeated string designs = 18;
     */
    designs: string[];
    /**
     * @gotags: dynamo:"FabricColors"
     *
     * @generated from protobuf field: repeated string colors = 19;
     */
    colors: string[];
    /**
     * @gotags: dynamo:"FabricOrigin"
     *
     * @generated from protobuf field: string origin = 20;
     */
    origin: string;
    /**
     * @gotags: dynamo:"FabricPillingGrade"
     *
     * @generated from protobuf field: string pilling_grade = 21;
     */
    pillingGrade: string;
    /**
     * @gotags: dynamo:"FabricFireCode"
     *
     * @generated from protobuf field: string fire_code = 22;
     */
    fireCode: string;
    /**
     * @gotags: dynamo:"FabricCategories"
     *
     * @generated from protobuf field: repeated string categories = 23;
     */
    categories: string[];
    /**
     * @gotags: dynamo:"FabricBrand"
     *
     * @generated from protobuf field: string brand = 24;
     */
    brand: string;
    /**
     * @gotags: dynamo:"FabricIsDropShipped"
     *
     * @generated from protobuf field: bool is_drop_shipped = 25;
     */
    isDropShipped: boolean;
    /**
     * @gotags: dynamo:"FabricMainImage"
     *
     * @generated from protobuf field: string image = 26;
     */
    image: string;
    /**
     * @gotags: dynamo:"FabricAPIDisplayPrice"
     *
     * @generated from protobuf field: string display_price = 27;
     */
    displayPrice: string;
    /**
     * @gotags: dynamo:"FabricAPIPrice"
     *
     * @generated from protobuf field: int32 price = 28;
     */
    price: number;
    /**
     * @gotags: dynamo:"FabricOtherImages"
     *
     * @generated from protobuf field: repeated string supplemental_images = 29;
     */
    supplementalImages: string[];
    /**
     * @gotags: dynamo:"FabricFullWidthImages"
     *
     * @generated from protobuf field: repeated string full_width_images = 30;
     */
    fullWidthImages: string[];
    /**
     * @gotags: dynamo:"-"
     *
     * @generated from protobuf field: magnoliafabrics.Inventory inventory = 31;
     */
    inventory?: Inventory;
}
// @generated message type with reflection information, may provide speed optimized methods
class Fabric$Type extends MessageType<Fabric> {
    constructor() {
        super("magnoliafabrics.Fabric", [
            { no: 1, name: "sku", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "product_code", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "pattern", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "color", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "pattern_color_combo", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 6, name: "status", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 7, name: "intro_date", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 8, name: "contents", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 9, name: "width", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 10, name: "direction", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 11, name: "h_repeat", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 12, name: "v_repeat", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 13, name: "cleaning_code", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 14, name: "double_rubs", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 15, name: "misc", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 16, name: "disclaimer", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 17, name: "uses", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 18, name: "designs", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 19, name: "colors", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 20, name: "origin", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 21, name: "pilling_grade", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 22, name: "fire_code", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 23, name: "categories", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 24, name: "brand", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 25, name: "is_drop_shipped", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 26, name: "image", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 27, name: "display_price", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 28, name: "price", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 29, name: "supplemental_images", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 30, name: "full_width_images", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 31, name: "inventory", kind: "message", T: () => Inventory }
        ]);
    }
    create(value?: PartialMessage<Fabric>): Fabric {
        const message = { sku: "", productCode: "", pattern: "", color: "", patternColorCombo: "", status: "", introDate: "", contents: "", width: "", direction: "", hRepeat: "", vRepeat: "", cleaningCode: "", doubleRubs: "", misc: "", disclaimer: "", uses: [], designs: [], colors: [], origin: "", pillingGrade: "", fireCode: "", categories: [], brand: "", isDropShipped: false, image: "", displayPrice: "", price: 0, supplementalImages: [], fullWidthImages: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Fabric>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Fabric): Fabric {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string sku */ 1:
                    message.sku = reader.string();
                    break;
                case /* string product_code */ 2:
                    message.productCode = reader.string();
                    break;
                case /* string pattern */ 3:
                    message.pattern = reader.string();
                    break;
                case /* string color */ 4:
                    message.color = reader.string();
                    break;
                case /* string pattern_color_combo */ 5:
                    message.patternColorCombo = reader.string();
                    break;
                case /* string status */ 6:
                    message.status = reader.string();
                    break;
                case /* string intro_date */ 7:
                    message.introDate = reader.string();
                    break;
                case /* string contents */ 8:
                    message.contents = reader.string();
                    break;
                case /* string width */ 9:
                    message.width = reader.string();
                    break;
                case /* string direction */ 10:
                    message.direction = reader.string();
                    break;
                case /* string h_repeat */ 11:
                    message.hRepeat = reader.string();
                    break;
                case /* string v_repeat */ 12:
                    message.vRepeat = reader.string();
                    break;
                case /* string cleaning_code */ 13:
                    message.cleaningCode = reader.string();
                    break;
                case /* string double_rubs */ 14:
                    message.doubleRubs = reader.string();
                    break;
                case /* string misc */ 15:
                    message.misc = reader.string();
                    break;
                case /* string disclaimer */ 16:
                    message.disclaimer = reader.string();
                    break;
                case /* repeated string uses */ 17:
                    message.uses.push(reader.string());
                    break;
                case /* repeated string designs */ 18:
                    message.designs.push(reader.string());
                    break;
                case /* repeated string colors */ 19:
                    message.colors.push(reader.string());
                    break;
                case /* string origin */ 20:
                    message.origin = reader.string();
                    break;
                case /* string pilling_grade */ 21:
                    message.pillingGrade = reader.string();
                    break;
                case /* string fire_code */ 22:
                    message.fireCode = reader.string();
                    break;
                case /* repeated string categories */ 23:
                    message.categories.push(reader.string());
                    break;
                case /* string brand */ 24:
                    message.brand = reader.string();
                    break;
                case /* bool is_drop_shipped */ 25:
                    message.isDropShipped = reader.bool();
                    break;
                case /* string image */ 26:
                    message.image = reader.string();
                    break;
                case /* string display_price */ 27:
                    message.displayPrice = reader.string();
                    break;
                case /* int32 price */ 28:
                    message.price = reader.int32();
                    break;
                case /* repeated string supplemental_images */ 29:
                    message.supplementalImages.push(reader.string());
                    break;
                case /* repeated string full_width_images */ 30:
                    message.fullWidthImages.push(reader.string());
                    break;
                case /* magnoliafabrics.Inventory inventory */ 31:
                    message.inventory = Inventory.internalBinaryRead(reader, reader.uint32(), options, message.inventory);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Fabric, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string sku = 1; */
        if (message.sku !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.sku);
        /* string product_code = 2; */
        if (message.productCode !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.productCode);
        /* string pattern = 3; */
        if (message.pattern !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.pattern);
        /* string color = 4; */
        if (message.color !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.color);
        /* string pattern_color_combo = 5; */
        if (message.patternColorCombo !== "")
            writer.tag(5, WireType.LengthDelimited).string(message.patternColorCombo);
        /* string status = 6; */
        if (message.status !== "")
            writer.tag(6, WireType.LengthDelimited).string(message.status);
        /* string intro_date = 7; */
        if (message.introDate !== "")
            writer.tag(7, WireType.LengthDelimited).string(message.introDate);
        /* string contents = 8; */
        if (message.contents !== "")
            writer.tag(8, WireType.LengthDelimited).string(message.contents);
        /* string width = 9; */
        if (message.width !== "")
            writer.tag(9, WireType.LengthDelimited).string(message.width);
        /* string direction = 10; */
        if (message.direction !== "")
            writer.tag(10, WireType.LengthDelimited).string(message.direction);
        /* string h_repeat = 11; */
        if (message.hRepeat !== "")
            writer.tag(11, WireType.LengthDelimited).string(message.hRepeat);
        /* string v_repeat = 12; */
        if (message.vRepeat !== "")
            writer.tag(12, WireType.LengthDelimited).string(message.vRepeat);
        /* string cleaning_code = 13; */
        if (message.cleaningCode !== "")
            writer.tag(13, WireType.LengthDelimited).string(message.cleaningCode);
        /* string double_rubs = 14; */
        if (message.doubleRubs !== "")
            writer.tag(14, WireType.LengthDelimited).string(message.doubleRubs);
        /* string misc = 15; */
        if (message.misc !== "")
            writer.tag(15, WireType.LengthDelimited).string(message.misc);
        /* string disclaimer = 16; */
        if (message.disclaimer !== "")
            writer.tag(16, WireType.LengthDelimited).string(message.disclaimer);
        /* repeated string uses = 17; */
        for (let i = 0; i < message.uses.length; i++)
            writer.tag(17, WireType.LengthDelimited).string(message.uses[i]);
        /* repeated string designs = 18; */
        for (let i = 0; i < message.designs.length; i++)
            writer.tag(18, WireType.LengthDelimited).string(message.designs[i]);
        /* repeated string colors = 19; */
        for (let i = 0; i < message.colors.length; i++)
            writer.tag(19, WireType.LengthDelimited).string(message.colors[i]);
        /* string origin = 20; */
        if (message.origin !== "")
            writer.tag(20, WireType.LengthDelimited).string(message.origin);
        /* string pilling_grade = 21; */
        if (message.pillingGrade !== "")
            writer.tag(21, WireType.LengthDelimited).string(message.pillingGrade);
        /* string fire_code = 22; */
        if (message.fireCode !== "")
            writer.tag(22, WireType.LengthDelimited).string(message.fireCode);
        /* repeated string categories = 23; */
        for (let i = 0; i < message.categories.length; i++)
            writer.tag(23, WireType.LengthDelimited).string(message.categories[i]);
        /* string brand = 24; */
        if (message.brand !== "")
            writer.tag(24, WireType.LengthDelimited).string(message.brand);
        /* bool is_drop_shipped = 25; */
        if (message.isDropShipped !== false)
            writer.tag(25, WireType.Varint).bool(message.isDropShipped);
        /* string image = 26; */
        if (message.image !== "")
            writer.tag(26, WireType.LengthDelimited).string(message.image);
        /* string display_price = 27; */
        if (message.displayPrice !== "")
            writer.tag(27, WireType.LengthDelimited).string(message.displayPrice);
        /* int32 price = 28; */
        if (message.price !== 0)
            writer.tag(28, WireType.Varint).int32(message.price);
        /* repeated string supplemental_images = 29; */
        for (let i = 0; i < message.supplementalImages.length; i++)
            writer.tag(29, WireType.LengthDelimited).string(message.supplementalImages[i]);
        /* repeated string full_width_images = 30; */
        for (let i = 0; i < message.fullWidthImages.length; i++)
            writer.tag(30, WireType.LengthDelimited).string(message.fullWidthImages[i]);
        /* magnoliafabrics.Inventory inventory = 31; */
        if (message.inventory)
            Inventory.internalBinaryWrite(message.inventory, writer.tag(31, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message magnoliafabrics.Fabric
 */
export const Fabric = new Fabric$Type();
