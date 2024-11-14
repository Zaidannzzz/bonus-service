"use client";

import Paper from "@mui/material/Paper";
import React from "react";

const PaperCard: React.FC<{
  children: React.ReactNode;
  sx?: object;
  props?: object;
}> = ({ children, sx, ...props }) => {
  return (
    <Paper
      sx={{
        backgroundColor: "#FFFFFF",
        px: 2.5,
        py: 2,
        outline: `1px solid #d9d9d9`,
        borderRadius: "4px",
        boxShadow: "none",
        ...sx,
      }}
      {...props}
    >
      {children}
    </Paper>
  );
};

export default PaperCard;
