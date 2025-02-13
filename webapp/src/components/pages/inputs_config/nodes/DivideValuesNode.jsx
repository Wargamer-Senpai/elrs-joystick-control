// SPDX-FileCopyrightText: © 2023 OneEyeFPV oneeyefpv@gmail.com
// SPDX-License-Identifier: GPL-3.0-or-later
// SPDX-License-Identifier: FS-0.9-or-later

import React from 'react';

import {SvgIcon} from "@mui/material";
import {TopOutput} from "../handles/TopOutput";

import {TruthyInputHandle} from "../handles/TruthyInputHandle";
import {FalsyInputHandle} from "../handles/FalsyInputHandle";
import {GenericInputNode} from "./GenericInputNode";

function DivideValuesNode(node) {

    return (<GenericInputNode
        node={node}
        valueProps={{
            style: {marginTop: "0px", marginBottom: "2px"}
        }}
        iconProps={{
            style: {marginBottom: "-6px"}
        }}
        labelProps={{
            style: {marginTop: "2px", marginBottom: "0px"}
        }}
    >
        <TopOutput node={node}/>


        <TruthyInputHandle node={node} fieldName={"left"}/>
        <FalsyInputHandle node={node} fieldName={"right"}/>
    </GenericInputNode>);
}

DivideValuesNode.type = "divide";
DivideValuesNode.menuIcon = <SvgIcon>
    <g>
        <line x1="4" y1="12" x2="20" y2="12"
              fill="none"
              stroke="#656565"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
        />
        <circle cx="12" cy="6" r="1.5" fill="#656565"/>
        <circle cx="12" cy="18" r="1.5" fill="#656565"/>
    </g>
</SvgIcon>;

export default DivideValuesNode;
