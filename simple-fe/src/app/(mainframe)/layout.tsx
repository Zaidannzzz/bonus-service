"use client";

import React from "react";

import { Box, Stack } from "@mui/material";
import Header from "./header/header";

const MainframeLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <Stack>
      <Header />
      <Box sx={{ display: "flex", overflowY: "auto" }}>
        <Box
          sx={{
            flexGrow: 1,
            px: 3,
            py: 2,
            maxHeight: `calc(100vh - 96px)`,
          }}
        >
          {children}
        </Box>
      </Box>
    </Stack>
  );
};

export default MainframeLayout;
