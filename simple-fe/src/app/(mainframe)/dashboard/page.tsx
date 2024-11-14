"use client";

import React, { useState } from "react";
import { Box, Button, Skeleton, Stack, Typography } from "@mui/material";
import { useAtom } from "jotai";
import { RESET } from "jotai/utils";
import { useRouter } from "next/navigation";
import toast, { Toaster } from "react-hot-toast";
import { userAtom } from "@/app/atoms/user";
import PaperCard from "@/app/components/paperCard";
import { isEmpty } from "lodash";

const Page = () => {
  const [user, setUser] = useAtom(userAtom);
  const [mounted, setMounted] = useState<boolean>(false);
  const router = useRouter();

  React.useEffect(() => {
    setMounted(true);
  }, []);

  const handleLogout = async () => {
    try {
      setUser(RESET);
      localStorage.clear();
      router.replace(`/login`);
      toast.success("Sukses logout", { id: "logout" });
    } catch (_) {
      console.log(_);
      toast.error("Gagal logout", { id: "logout" });
    }
  };

  const profileContent = React.useMemo(() => {
    if (!mounted) return [];
    return [
      { title: "id", value: !isEmpty(user?.uuid) ? user?.uuid : "-" },
      { title: "Nama", value: !isEmpty(user?.name) ? user?.name : "-" },
      { title: "Email", value: !isEmpty(user?.email) ? user?.email : "-" },
      { title: "Gender", value: !isEmpty(user?.gender) ? user?.gender : "-" },
    ];
  }, [user, mounted]);

  return (
    <Stack spacing={2} direction="row" justifyContent="center" alignItems="center" sx={{ height: "70vh" }}>
      <PaperCard sx={{ width: { md: "40%" }, p: 2 }}>
        <Typography variant="h5" color="#0F4C82" sx={{ mb: 2 }}>
          Profile
        </Typography>

        <Stack spacing={3}>
          {/* just that lazy :) */}
          {isEmpty(profileContent) && (
            <>
              <Skeleton animation="wave" />
              <Skeleton animation="wave" />
              <Skeleton animation="wave" />
              <Skeleton animation="wave" />
              <Skeleton animation="wave" />
              <Skeleton animation="wave" />
              <Skeleton animation="wave" />
            </>
          )}
          {profileContent?.map((data, index) => (
            <Box key={index}>
              <Typography variant="h6" color="#0F4C82">
                {data.title}
              </Typography>
              <Typography variant="body1" color="#666666">
                {data.value}
              </Typography>
            </Box>
          ))}
        </Stack>

        <Button variant="contained" size="large" onClick={handleLogout} color="error" sx={{ mt: 3 }}>
          Logout
        </Button>
      </PaperCard>
      <Toaster />
    </Stack>
  );
};
export default Page;
