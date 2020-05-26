import React from "react";
import { TenantStatus } from "@teamdream/models";
import { Message } from "./Message";
import { useTeamdream } from "@teamdream/hooks";

export const TenantStatusInfo = () => {
  const teamdream = useTeamdream();

  if (!teamdream.isBillingEnabled() || teamdream.session.tenant.status !== TenantStatus.Locked) {
    return null;
  }

  return (
    <div className="container">
      <Message type="error">
        This site is locked due to lack of a subscription. Visit the <a href="/admin/billing">Billing</a> settings to
        update it.
      </Message>
    </div>
  );
};
