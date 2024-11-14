"use client";

import React, { useEffect } from "react";
import { useRouter } from "next/navigation";
import { InputAdornment, Stack, TextField, Typography, useTheme } from "@mui/material";
import { LoadingButton } from "@mui/lab";
import { IconEye, IconEyeOff, IconLogin2 } from "@tabler/icons-react";
import { useAtom } from "jotai";
import toast, { Toaster } from "react-hot-toast";

import { RESET } from "jotai/utils";
import { userAtom } from "@/app/atoms/user";
import authenticationV1 from "@/app/services/authentication";
import { tokenAtom } from "@/app/atoms/token";
import axios from "axios";

const LoginPage = () => {
  const theme = useTheme();
  const router = useRouter();
  const [, setUser] = useAtom(userAtom);
  const [, setToken] = useAtom(tokenAtom);

  const [loading, setLoading] = React.useState(false);
  const [showPassword, setShowPassword] = React.useState(false);
  const [inputValue, setInputValue] = React.useState({
    email: "",
    password: "",
  });

  const handleKeyDown = (source: string, e: React.KeyboardEvent) => {
    if (e.key === "Enter" && source === "password") {
      login();
    }
  };

  const login = async () => {
    setLoading(true);
    try {
      const accessToken = await authenticationV1.login(inputValue.email, inputValue.password);

      setToken({ value: accessToken, id: null });
      toast.success("Sukses login", { id: "login" });
      router.push("/dashboard");
    } catch (error: unknown) {
      if (axios.isAxiosError(error)) {
        toast.error(`Login Gagal, ${error.response?.data?.message}`, { id: "login" });
      } else {
        toast.error("Login Gagal", { id: "login" });
      }
      router.push("/login");
    }

    setLoading(false);
  };

  useEffect(() => {
    setUser(RESET);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <Stack spacing={2} alignItems={"center"} width={"100%"}>
      <Typography variant="h6" fontWeight={600} color={theme.palette.primary.main}>
        Login
      </Typography>
      <TextField
        fullWidth
        size="small"
        label={"Email"}
        type={"email"}
        value={inputValue.email}
        onChange={(e) => setInputValue({ ...inputValue, email: e.target.value })}
        onKeyDown={(e: React.KeyboardEvent) => handleKeyDown("email", e)}
      />
      <TextField
        fullWidth
        size="small"
        label={"Password"}
        type={showPassword ? "text" : "password"}
        value={inputValue.password}
        onChange={(e) => setInputValue({ ...inputValue, password: e.target.value })}
        InputProps={{
          endAdornment: (
            <InputAdornment position="end" sx={{ cursor: "pointer" }}>
              {showPassword ? (
                <IconEyeOff onClick={() => setShowPassword(false)} />
              ) : (
                <IconEye onClick={() => setShowPassword(true)} />
              )}
            </InputAdornment>
          ),
        }}
        onKeyDown={(e: React.KeyboardEvent) => handleKeyDown("password", e)}
      />
      <LoadingButton
        fullWidth
        loading={loading}
        disabled={!inputValue.email || !inputValue.password}
        variant="contained"
        onClick={() => login()}
        startIcon={<IconLogin2 size={16} />}
      >
        <Typography variant="body2" fontWeight={600} textTransform={"none"}>
          Login
        </Typography>
      </LoadingButton>
      <Toaster />
    </Stack>
  );
};

export default LoginPage;
