import { atomWithStorage } from "jotai/utils";

export const userAtom = atomWithStorage("user", {
  id: null,
  name: "",
  uuid: "",
  gender: "",
  email: "",
});
