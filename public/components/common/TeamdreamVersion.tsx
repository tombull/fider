import React from "react";
import { useTeamdream } from "@teamdream/hooks";

export const TeamdreamVersion = () => {
  const teamdream = useTeamdream();

  return <p className="info center hidden-sm hidden-md">Teamdream v{teamdream.settings.version}</p>;
};
