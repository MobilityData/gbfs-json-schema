# Typescript GBFS Language Bindings

[![npm version](https://badge.fury.io/js/gbfs-typescript-types.svg)](http://badge.fury.io/js/gbfs-typescript-types)

TypeScript types for parsing and working with General Bikeshare Feed Specification (GBFS) data, ensuring type safety and code consistency in TypeScript projects.

## Add the Dependency

To use `gbfs-typescript-types` in your own project, you need to
first install our [Node.js npm package](https://www.npmjs.com/package/gbfs-typescript-types):

```
npm install gbfs-typescript-types --save-dev
```

## Versions
- v2.3
- v3.0
- v3.1-RC

## Example Code
```typescript
// top level import
import { v3 } from 'gbfs-typescript-types';

// high level types
import { SystemInformation } from 'gbfs-typescript-types/v3.0';

// lower level properties need to have path specified
import { Data as SystemInformationData } from 'gbfs-typescript-types/v3.0/system_information';
import { Data as VehicleStatusData } from 'gbfs-typescript-types/v3.0/vehicle_status';

let vehicleStatus: v3.VehicleStatus;
let systemInformationData: SystemInformationData;
const url = "https://berlin.example.tier-services.io/tier_paris/gbfs/3.0/system-information";
fetch(url).then((systemInformationResponse) => {
    systemInformationResponse.json().then((systemInformationJson: SystemInformation) => {
        systemInformationJson // will have access to types
        systemInformationData = systemInformationJson.data;
    })
})
```
