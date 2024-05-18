// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { GetAuditUserArgs, GetAuditUserResult, GetAuditUserOutputArgs } from "./getAuditUser";
export const getAuditUser: typeof import("./getAuditUser").getAuditUser = null as any;
export const getAuditUserOutput: typeof import("./getAuditUser").getAuditUserOutput = null as any;
utilities.lazyLoad(exports, ["getAuditUser","getAuditUserOutput"], () => require("./getAuditUser"));

export { GetGroupArgs, GetGroupResult, GetGroupOutputArgs } from "./getGroup";
export const getGroup: typeof import("./getGroup").getGroup = null as any;
export const getGroupOutput: typeof import("./getGroup").getGroupOutput = null as any;
utilities.lazyLoad(exports, ["getGroup","getGroupOutput"], () => require("./getGroup"));

export { GetRetrieveMemberArgs, GetRetrieveMemberResult, GetRetrieveMemberOutputArgs } from "./getRetrieveMember";
export const getRetrieveMember: typeof import("./getRetrieveMember").getRetrieveMember = null as any;
export const getRetrieveMemberOutput: typeof import("./getRetrieveMember").getRetrieveMemberOutput = null as any;
utilities.lazyLoad(exports, ["getRetrieveMember","getRetrieveMemberOutput"], () => require("./getRetrieveMember"));

export { GetServiceAccountArgs, GetServiceAccountResult, GetServiceAccountOutputArgs } from "./getServiceAccount";
export const getServiceAccount: typeof import("./getServiceAccount").getServiceAccount = null as any;
export const getServiceAccountOutput: typeof import("./getServiceAccount").getServiceAccountOutput = null as any;
utilities.lazyLoad(exports, ["getServiceAccount","getServiceAccountOutput"], () => require("./getServiceAccount"));

export { GetServiceAccountTokenArgs, GetServiceAccountTokenResult, GetServiceAccountTokenOutputArgs } from "./getServiceAccountToken";
export const getServiceAccountToken: typeof import("./getServiceAccountToken").getServiceAccountToken = null as any;
export const getServiceAccountTokenOutput: typeof import("./getServiceAccountToken").getServiceAccountTokenOutput = null as any;
utilities.lazyLoad(exports, ["getServiceAccountToken","getServiceAccountTokenOutput"], () => require("./getServiceAccountToken"));

export { GetUserArgs, GetUserResult, GetUserOutputArgs } from "./getUser";
export const getUser: typeof import("./getUser").getUser = null as any;
export const getUserOutput: typeof import("./getUser").getUserOutput = null as any;
utilities.lazyLoad(exports, ["getUser","getUserOutput"], () => require("./getUser"));

export { GetWorkplaceArgs, GetWorkplaceResult } from "./getWorkplace";
export const getWorkplace: typeof import("./getWorkplace").getWorkplace = null as any;
export const getWorkplaceOutput: typeof import("./getWorkplace").getWorkplaceOutput = null as any;
utilities.lazyLoad(exports, ["getWorkplace","getWorkplaceOutput"], () => require("./getWorkplace"));

export { GetWorkplaceRoleArgs, GetWorkplaceRoleResult, GetWorkplaceRoleOutputArgs } from "./getWorkplaceRole";
export const getWorkplaceRole: typeof import("./getWorkplaceRole").getWorkplaceRole = null as any;
export const getWorkplaceRoleOutput: typeof import("./getWorkplaceRole").getWorkplaceRoleOutput = null as any;
utilities.lazyLoad(exports, ["getWorkplaceRole","getWorkplaceRoleOutput"], () => require("./getWorkplaceRole"));

export { GroupsArgs } from "./groups";
export type Groups = import("./groups").Groups;
export const Groups: typeof import("./groups").Groups = null as any;
utilities.lazyLoad(exports, ["Groups"], () => require("./groups"));

export { GroupsMemberArgs } from "./groupsMember";
export type GroupsMember = import("./groupsMember").GroupsMember;
export const GroupsMember: typeof import("./groupsMember").GroupsMember = null as any;
utilities.lazyLoad(exports, ["GroupsMember"], () => require("./groupsMember"));

export { ListGroupsArgs, ListGroupsResult } from "./listGroups";
export const listGroups: typeof import("./listGroups").listGroups = null as any;
export const listGroupsOutput: typeof import("./listGroups").listGroupsOutput = null as any;
utilities.lazyLoad(exports, ["listGroups","listGroupsOutput"], () => require("./listGroups"));

export { ListInvitesArgs, ListInvitesResult } from "./listInvites";
export const listInvites: typeof import("./listInvites").listInvites = null as any;
export const listInvitesOutput: typeof import("./listInvites").listInvitesOutput = null as any;
utilities.lazyLoad(exports, ["listInvites","listInvitesOutput"], () => require("./listInvites"));

export { ListServiceAccountTokensArgs, ListServiceAccountTokensResult, ListServiceAccountTokensOutputArgs } from "./listServiceAccountTokens";
export const listServiceAccountTokens: typeof import("./listServiceAccountTokens").listServiceAccountTokens = null as any;
export const listServiceAccountTokensOutput: typeof import("./listServiceAccountTokens").listServiceAccountTokensOutput = null as any;
utilities.lazyLoad(exports, ["listServiceAccountTokens","listServiceAccountTokensOutput"], () => require("./listServiceAccountTokens"));

export { ListServiceAccountsArgs, ListServiceAccountsResult } from "./listServiceAccounts";
export const listServiceAccounts: typeof import("./listServiceAccounts").listServiceAccounts = null as any;
export const listServiceAccountsOutput: typeof import("./listServiceAccounts").listServiceAccountsOutput = null as any;
utilities.lazyLoad(exports, ["listServiceAccounts","listServiceAccountsOutput"], () => require("./listServiceAccounts"));

export { ListUsersArgs, ListUsersResult } from "./listUsers";
export const listUsers: typeof import("./listUsers").listUsers = null as any;
export const listUsersOutput: typeof import("./listUsers").listUsersOutput = null as any;
utilities.lazyLoad(exports, ["listUsers","listUsersOutput"], () => require("./listUsers"));

export { ListWorkplaceRolesArgs, ListWorkplaceRolesResult } from "./listWorkplaceRoles";
export const listWorkplaceRoles: typeof import("./listWorkplaceRoles").listWorkplaceRoles = null as any;
export const listWorkplaceRolesOutput: typeof import("./listWorkplaceRoles").listWorkplaceRolesOutput = null as any;
utilities.lazyLoad(exports, ["listWorkplaceRoles","listWorkplaceRolesOutput"], () => require("./listWorkplaceRoles"));

export { ListWorkplaceRolesPermissionsArgs, ListWorkplaceRolesPermissionsResult } from "./listWorkplaceRolesPermissions";
export const listWorkplaceRolesPermissions: typeof import("./listWorkplaceRolesPermissions").listWorkplaceRolesPermissions = null as any;
export const listWorkplaceRolesPermissionsOutput: typeof import("./listWorkplaceRolesPermissions").listWorkplaceRolesPermissionsOutput = null as any;
utilities.lazyLoad(exports, ["listWorkplaceRolesPermissions","listWorkplaceRolesPermissionsOutput"], () => require("./listWorkplaceRolesPermissions"));

export { ServiceAccountTokensArgs } from "./serviceAccountTokens";
export type ServiceAccountTokens = import("./serviceAccountTokens").ServiceAccountTokens;
export const ServiceAccountTokens: typeof import("./serviceAccountTokens").ServiceAccountTokens = null as any;
utilities.lazyLoad(exports, ["ServiceAccountTokens"], () => require("./serviceAccountTokens"));

export { ServiceAccountsArgs } from "./serviceAccounts";
export type ServiceAccounts = import("./serviceAccounts").ServiceAccounts;
export const ServiceAccounts: typeof import("./serviceAccounts").ServiceAccounts = null as any;
utilities.lazyLoad(exports, ["ServiceAccounts"], () => require("./serviceAccounts"));

export { WorkplaceRolesArgs } from "./workplaceRoles";
export type WorkplaceRoles = import("./workplaceRoles").WorkplaceRoles;
export const WorkplaceRoles: typeof import("./workplaceRoles").WorkplaceRoles = null as any;
utilities.lazyLoad(exports, ["WorkplaceRoles"], () => require("./workplaceRoles"));

export { WorkplaceUpdateArgs } from "./workplaceUpdate";
export type WorkplaceUpdate = import("./workplaceUpdate").WorkplaceUpdate;
export const WorkplaceUpdate: typeof import("./workplaceUpdate").WorkplaceUpdate = null as any;
utilities.lazyLoad(exports, ["WorkplaceUpdate"], () => require("./workplaceUpdate"));


// Export enums:
export * from "../../types/enums/workplace/v3";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "doppler-native:workplace/v3:Groups":
                return new Groups(name, <any>undefined, { urn })
            case "doppler-native:workplace/v3:GroupsMember":
                return new GroupsMember(name, <any>undefined, { urn })
            case "doppler-native:workplace/v3:ServiceAccountTokens":
                return new ServiceAccountTokens(name, <any>undefined, { urn })
            case "doppler-native:workplace/v3:ServiceAccounts":
                return new ServiceAccounts(name, <any>undefined, { urn })
            case "doppler-native:workplace/v3:WorkplaceRoles":
                return new WorkplaceRoles(name, <any>undefined, { urn })
            case "doppler-native:workplace/v3:WorkplaceUpdate":
                return new WorkplaceUpdate(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("doppler-native", "workplace/v3", _module)
