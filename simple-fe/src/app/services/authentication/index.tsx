import axiosClient from "../axios";

const login = async (email: string, password: string) => {
  const payload = {
    email: email,
    password: password,
  };
  const data = await axiosClient.post(`${process.env.NEXT_PUBLIC_LOGIN_URL}`, payload);
  return data?.data?.data?.access_token;
};

const getUserProfile = async (token: string) => {
  return await axiosClient.get(`${process.env.NEXT_PUBLIC_GET_PROFILE_URL}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};

const authenticationV1 = {
  login,
  getUserProfile,
};

export default authenticationV1;
