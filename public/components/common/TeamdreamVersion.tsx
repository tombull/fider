import React from "react";
import { useTeamdream } from "@teamdream/hooks";

export const TeamdreamVersion = () => {
  const teamdream = useTeamdream();

  return (
    <p className="info center hidden-sm hidden-md">
      {!teamdream.isBillingEnabled() && (
        <>
          Support our{" "}
          <a target="_blank" href="http://opencollective.com/teamdream">
            OpenCollective
          </a>
          <br />
        </>
      )}
      Teamdream v{teamdream.settings.version}
    </p>
  );
};
