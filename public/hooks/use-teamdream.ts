import { useContext } from "react";
import { TeamdreamContext } from "@teamdream/services";

export const useTeamdream = () => useContext(TeamdreamContext);
