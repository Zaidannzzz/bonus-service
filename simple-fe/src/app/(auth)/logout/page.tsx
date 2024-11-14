"use client";

import { userAtom } from "@/app/atoms/user";
import { Stack, Typography } from "@mui/material";
import { useAtom } from "jotai";
import { RESET } from "jotai/utils";
import { useRouter } from "next/navigation";
import React from "react";
import toast from "react-hot-toast";

const LogoutPage = () => {
  const router = useRouter();

  const [mounted, setMounted] = React.useState<boolean>(false);
  const [, setUser] = useAtom(userAtom);

  const logout = async () => {
    try {
      setUser(RESET);

      router.replace("/login");
    } catch (_) {
      console.log(_);
      toast.error("Gagal logout", { id: "logout" });
    }
  };

  React.useEffect(() => {
    setMounted(true);
  }, []);

  React.useEffect(() => {
    if (!mounted) return;

    logout();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [mounted]);

  return (
    <Stack spacing={2}>
      <Typography>Logging out ...</Typography>
    </Stack>
  );
};

export default LogoutPage;
