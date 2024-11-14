"use client";

import { tokenAtom } from "@/app/atoms/token";
import { userAtom } from "@/app/atoms/user";
import authenticationV1 from "@/app/services/authentication";
import { Avatar, Box, Grid2, Stack, Typography } from "@mui/material";
import { useAtom } from "jotai";
import { useRouter } from "next/navigation";
import React, { useEffect, useState } from "react";
import toast from "react-hot-toast";

const Header = () => {
  const router = useRouter();
  const [, setUser] = useAtom(userAtom);
  const [token] = useAtom(tokenAtom);
  const [isMounted, setIsMounted] = useState<boolean>(false);

  const getProfile = async () => {
    try {
      const response = await authenticationV1.getUserProfile(token.value);

      setUser(response?.data?.data);
    } catch (error) {
      console.log("🚀 ~ getProfile ~ error:", error);
      toast.error("Session expired. Please log in to continue.");
      toast.error(`Login Gagal, ${error ?? ""}`, { id: "auth" });
      router.replace("/login");
    }
  };

  useEffect(() => {
    if (!isMounted) return;

    getProfile();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [isMounted]);

  useEffect(() => setIsMounted(true), []);

  return (
    <Box bgcolor={"#0F4C82"} sx={{ px: 3, zIndex: (theme) => theme.zIndex.drawer + 1 }}>
      <Box minHeight={64} display={"flex"} justifyContent={"space-between"} alignItems={"center"}>
        <Box
          display={"flex"}
          gap={1}
          alignItems={"center"}
          onClick={() => router.push("/dashboard")}
          sx={{ cursor: "pointer" }}
        >
          <Grid2 container spacing={2} justifyContent={"space-between"} alignItems={"center"}>
            <Grid2>
              <Stack>
                <Typography variant="body1" fontWeight={600} color={"white"}>
                  Devops Test
                </Typography>
                <Typography variant="body2" color={"white"}>
                  v1.0.0
                </Typography>
              </Stack>
            </Grid2>
          </Grid2>
        </Box>

        <Box display={"flex"} gap={2} alignItems={"center"} onClick={() => {}} sx={{ cursor: "pointer" }}>
          <Stack alignItems={"flex-end"}>
            <Typography color={"#F8F8FF"}>Alvindo</Typography>
            <Typography color={"#F8F8FF"}>Raka</Typography>
          </Stack>
          <Avatar sx={{ width: 32, height: 32 }} />
        </Box>
      </Box>
    </Box>
  );
};

export default Header;
