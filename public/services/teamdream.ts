import { createContext } from "react";
import { CurrentUser, SystemSettings, Tenant } from "@teamdream/models";

export class TeamdreamSession {
  private pContextID: string;
  private pTenant: Tenant;
  private pUser: CurrentUser | undefined;
  private pProps: { [key: string]: any } = {};

  constructor(data: any) {
    this.pContextID = data.contextID;
    this.pProps = data.props;
    this.pUser = data.user;
    this.pTenant = data.tenant;
  }

  public get contextID(): string {
    return this.pContextID;
  }

  public get user(): CurrentUser {
    return this.pUser!;
  }

  public get tenant(): Tenant {
    return this.pTenant;
  }

  public get props(): { [key: string]: any } {
    return this.pProps;
  }

  public get isAuthenticated(): boolean {
    return !!this.pUser;
  }
}

export class TeamdreamImpl {
  private pSettings!: SystemSettings;
  private pSession!: TeamdreamSession;

  public initialize = (): TeamdreamImpl => {
    const el = document.getElementById("server-data");
    const data = el ? JSON.parse(el.textContent || el.innerText) : {};
    this.pSettings = data.settings;
    this.pSession = new TeamdreamSession(data);
    return this;
  };

  public get session(): TeamdreamSession {
    return this.pSession;
  }

  public get settings(): SystemSettings {
    return this.pSettings;
  }

  public isBillingEnabled(): boolean {
    return !!this.pSettings.stripePublicKey;
  }

  public isProduction(): boolean {
    return this.pSettings.environment === "production";
  }

  public isSingleHostMode(): boolean {
    return this.pSettings.mode === "single";
  }
}

export let Teamdream = new TeamdreamImpl();

export const TeamdreamContext = createContext<TeamdreamImpl>(Teamdream);
