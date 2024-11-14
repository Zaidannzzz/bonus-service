import { atomWithStorage } from "jotai/utils";

export const tokenAtom = atomWithStorage("token", {
  id: null,
  value: "",
});
