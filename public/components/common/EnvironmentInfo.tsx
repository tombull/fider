import React from "react";
import { useTeamdream } from "@teamdream/hooks";

export const EnvironmentInfo = () => {
  const teamdream = useTeamdream();

  if (teamdream.isProduction()) {
    return null;
  }

  return (
    <div className="c-env-info">
      Env: {teamdream.settings.environment} | Compiler: {teamdream.settings.compiler} | Version:{" "}
      {teamdream.settings.version} | BuildTime: {teamdream.settings.buildTime || "N/A"} |
      {!teamdream.isSingleHostMode() && `TenantID: ${teamdream.session.tenant.id}`} |{" "}
      {teamdream.session.isAuthenticated && `UserID: ${teamdream.session.user.id}`}
    </div>
  );
};
