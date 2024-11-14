import axios, { CreateAxiosDefaults } from "axios";

const clientConfig: CreateAxiosDefaults = {
  baseURL: process.env.NEXT_PUBLIC_BASE_API_URL,
  headers: {
    "Content-Type": "application/json; charset=UTF-8",
  },
};

const axiosClient = axios.create(clientConfig);

export default axiosClient;
