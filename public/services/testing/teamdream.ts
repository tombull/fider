import { Teamdream } from "@teamdream/services";

export const teamdreamMock = {
  notAuthenticated: () => {
    Teamdream.initialize();
    Object.defineProperty(Teamdream.session, "isAuthenticated", {
      get() {
        return false;
      },
    });
  },
  authenticated: () => {
    Teamdream.initialize();
    Object.defineProperty(Teamdream.session, "isAuthenticated", {
      get() {
        return true;
      },
    });
  },
};
