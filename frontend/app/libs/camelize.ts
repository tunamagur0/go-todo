interface StringObject {
  [key: string]: any;
}

export const camelize = (obj: Object): StringObject => {
  const ret: StringObject = {};
  for (const [key, value] of Object.entries(obj)) {
    ret[snakeToCamel(key)] = value;
  }
  return ret;
};

const snakeToCamel = (snake: string): string => {
  return snake.replace(/([_][a-z])/g, (group) =>
    group.toUpperCase().replace('_', '')
  );
};
