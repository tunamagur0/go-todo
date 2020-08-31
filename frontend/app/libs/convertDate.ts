interface StringObject {
  [key: string]: any;
}

export const convertDate = (obj: Object): StringObject => {
  const ret: StringObject = {};
  for (const [key, value] of Object.entries(obj)) {
    if ((key.includes('At') || key.includes('_at')) && value !== null) {
      ret[key] = new Date(value);
    } else {
      ret[key] = value;
    }
  }
  return ret;
};
