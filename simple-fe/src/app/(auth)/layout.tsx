import { Box, Container, Stack } from "@mui/material";
import React from "react";
import PaperCard from "../components/paperCard";

const AuthLayout: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  return (
    <Container fixed maxWidth={"xs"}>
      <Box px={2} display={"flex"} height={"100vh"} alignItems={"center"} justifyContent={"center"}>
        <Stack spacing={2} alignItems={"center"} sx={{ width: "100%" }}>
          <PaperCard sx={{ width: "100%", p: 2 }}>
            <Stack spacing={2} alignItems={"center"}>
              {children}
            </Stack>
          </PaperCard>
        </Stack>
      </Box>
    </Container>
  );
};

export default AuthLayout;
