// TODO this file is a travesty. It's also a utility script. Maybe come back later and fix it up...
// then again, maybe not

/**
 * Output:
 * {
 *  stringKey: "string"
 *  numberKey: "number"
 *  arrayKey: [{...}]
 *  objectKey: {...}
 *  required: []
 *  optional: []
 * }
 */

/**
 * @typedef {string | number | boolean  | JsonObject | JsonValueArray} JsonValue
 * @typedef {JsonValue[]} JsonValueArray
 * @typedef {{[key: string]: JsonValue}} JsonObject
 *
 * @typedef {"string" | "number" | "boolean"  | JsonSchema | JsonSchemaArray} JsonType
 * @typedef {JsonSchema[]} JsonSchemaArray
 *
 * @typedef {{
 *  [key: string]: JsonType | JsonSchema | Set<string>,
 *  required: Set<string>,
 *  optional: Set<string>,
 * }} JsonSchema
 */

/** @return {JsonSchema} */
function newJsonSchema() {
  return {
    required: new Set(),
    optional: new Set(),
  };
}

/**
 * @param {JsonSchema} schema
 * @param {string} property
 * @param {JsonType} value
 * @return {JsonSchema}
 */
function addToSchema(schema, property, value) {
  const singlePropertySchema = newJsonSchema();
  singlePropertySchema[property] = value;
  singlePropertySchema.required.add(property);

  unionSchema(schema, singlePropertySchema);
  return schema;
}

/** @type {{[key: string]: JsonSchema}} */
const schemas = {};
console.log("getSchemas:", function () {
  console.log(
    JSON.stringify(schemas, (_, value) =>
      value instanceof Set ? [...value] : value
    )
  );
});

/** @param {FileSystemFileHandle} fileHandle */
export async function documentJsonFile(fileHandle) {
  const file = await fileHandle.getFile();
  return documentJsonString(await file.text(), fileHandle.name);
}

/**
 * @param {string} json
 * @param {string} schemaKey
 */
export function documentJsonString(json, schemaKey) {
  console.log("documenting", schemaKey /* , json */);
  if (!schemas[schemaKey]) {
    schemas[schemaKey] = newJsonSchema();
  }
  const createdSchema = documentJson(JSON.parse(json), newJsonSchema());
  unionSchema(schemas[schemaKey], createdSchema);
}

/**
 * @param {JsonObject} json
 * @param {JsonSchema} currentSchema
 * @returns {JsonSchema}
 */
export function documentJson(json, currentSchema) {
  const newKeys = new Set(Object.keys(json));

  const currentSchemaHasKeys =
    currentSchema.optional.size || currentSchema.required.size;
  if (!currentSchemaHasKeys) {
    currentSchema.required = newKeys;
  }

  // keys that are in one set, but not in both
  const optionalKeys = currentSchema.required.symmetricDifference(newKeys);
  const requiredKeys = currentSchema.required.intersection(newKeys);

  for (const key of newKeys) {
    currentSchema[key] = documentValue(json[key]);
  }

  return currentSchema;
}

/**
 *
 * @param {JsonValue} value
 * @returns {JsonType}
 */
function documentValue(value) {
  switch (typeof value) {
    case "string":
      return "string";
    case "number":
      return "number";
    case "boolean":
      return "boolean";
    default:
      if (Array.isArray(value)) {
        return documentArray(value);
      } else /* Object */ {
        return documentJson(value, newJsonSchema());
      }
  }
}

/**
 * @param {any[]} array
 * @return {JsonSchema[]}
 */
function documentArray(array) {
  // TODO I think they all come out optional
  const schema = newJsonSchema();
  array.forEach((value) => {
    const singleSchema = documentJson(value, newJsonSchema());
    for (const [key, value] of Object.entries(singleSchema)) {
      if (!(value instanceof Set)) {
        // set properties are their own thing
        addToSchema(schema, key, value);
      }
    }
  });

  return [schema];
}

/**
 * @param {JsonSchema} destination
 * @param {JsonSchema} source
 * @returns {void}
 */
function unionSchema(destination, source) {
  // if destination is an empty schema
  if (isEmptySchema(destination)) {
    Object.assign(destination, source);
    return;
  }
  if (isEmptySchema(source)) {
    return;
  }

  // only keys that are required in both sets are actually required
  const newRequiredKeys = destination.required.intersection(source.required);
  const noLongerRequiredKeys = destination.required.symmetricDifference(
    source.required
  );
  const newOptionalKeys = destination.optional
    .union(source.optional)
    .union(noLongerRequiredKeys);

  // invert the intersection to get keys that were required but are now optional
  // merge the two sets of optional keys, plus the no-longer-required ones

  for (const key of source.required.union(source.optional)) {
    if (destination[key] === source[key]) continue;
    if (destination[key] === undefined) {
      destination[key] = source[key];
      continue;
    }
    if (Array.isArray(destination[key]) && Array.isArray(source[key])) {
      unionSchema(destination[key][0], source[key][0]);
      continue;
    }
    if (isJsonSchema(destination[key]) && isJsonSchema(source[key])) {
      unionSchema(destination[key], source[key]);
      continue;
    }
    console.log(key, "|", destination[key], source[key]);
    debugger;
    throw new Error("Unable to merge destination and source on: " + key);
  }
  destination.required = newRequiredKeys;
  destination.optional = newOptionalKeys;
}

/**
 * @param {JsonSchema[string]} value
 * @returns {value is JsonSchema}
 */
function isJsonSchema(value) {
  return typeof value === "object" && !Array.isArray(value);
}

/**
 * @param {JsonSchema} schema
 * @returns {boolean}
 */
function isEmptySchema(schema) {
  const propertyCount = Object.keys(schema).length;
  const noProperties = schema.required.size === 0 && schema.optional.size === 0;
  if (propertyCount === 2 && noProperties) return true;
  if (propertyCount > 2 && noProperties) debugger; // properties not logging correctly
  return false;
}
