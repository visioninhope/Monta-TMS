/*
 * COPYRIGHT(c) 2023 MONTA
 *
 * This file is part of Monta.
 *
 * Monta is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Monta is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Monta.  If not, see <https://www.gnu.org/licenses/>.
 */

/* eslint-disable react/jsx-no-target-blank */
import { AsideMenuItemWithSub } from "./AsideMenuItemWithSub";
import { AsideMenuItem } from "./AsideMenuItem";
import Image from "next/image";
import gen005 from "../../../../public/media/icons/duotune/general/gen005.svg";
import gen025 from "../../../../public/media/icons/duotune/general/gen025.svg";
import gen040 from "../../../../public/media/icons/duotune/general/gen040.svg";
import gen019 from "../../../../public/media/icons/duotune/general/gen025.svg";
import ecm007 from "../../../../public/media/icons/duotune/ecommerce/ecm007.svg";
import com006 from "../../../../public/media/icons/duotune/communication/com006.svg";
import gen022 from "../../../../public/media/icons/duotune/general/gen022.svg";
import com012 from "../../../../public/media/icons/duotune/communication/com012.svg";
import gen051 from "../../../../public/media/icons/duotune/general/gen051.svg";

export function AsideMenuMain() {

  return (
    <>
      <AsideMenuItem
        to="/dashboard"
        icon={gen025}
        title={"Dashboard"}
      />
      <AsideMenuItem
        to="/builder"
        icon={gen019}
        title="Layout Builder"
      />
      <div className="menu-item">
        <div className="menu-content pt-8 pb-2">
          <span className="menu-section text-muted text-uppercase fs-8 ls-1">Crafted</span>
        </div>
      </div>
      <AsideMenuItemWithSub
        to="/crafted/pages"
        title="Pages"
        icon={ecm007}
      >
        <AsideMenuItemWithSub to="/crafted/pages/profile" title="Profile" hasBullet={true}>
          <AsideMenuItem to="/crafted/pages/profile/overview" title="Overview" hasBullet={true} />
          <AsideMenuItem to="/crafted/pages/profile/projects" title="Projects" hasBullet={true} />
          <AsideMenuItem to="/crafted/pages/profile/campaigns" title="Campaigns" hasBullet={true} />
          <AsideMenuItem to="/crafted/pages/profile/documents" title="Documents" hasBullet={true} />
          <AsideMenuItem
            to="/crafted/pages/profile/connections"
            title="Connections"
            hasBullet={true}
          />
        </AsideMenuItemWithSub>

        <AsideMenuItemWithSub to="/crafted/pages/wizards" title="Wizards" hasBullet={true}>
          <AsideMenuItem
            to="/crafted/pages/wizards/horizontal"
            title="Horizontal"
            hasBullet={true}
          />
          <AsideMenuItem to="/crafted/pages/wizards/vertical" title="Vertical" hasBullet={true} />
        </AsideMenuItemWithSub>
      </AsideMenuItemWithSub>
      <AsideMenuItemWithSub
        to="/crafted/accounts"
        title="Accounts"
        icon={com006}
      >
        <AsideMenuItem to="/crafted/account/overview" title="Overview" hasBullet={true} />
        <AsideMenuItem to="/crafted/account/settings" title="Settings" hasBullet={true} />
      </AsideMenuItemWithSub>
      <AsideMenuItemWithSub
        to="/error"
        title="Errors"
        icon={gen040}
      >
        <AsideMenuItem to="/error/404" title="Error 404" hasBullet={true} />
        <AsideMenuItem to="/error/500" title="Error 500" hasBullet={true} />
      </AsideMenuItemWithSub>
      <AsideMenuItemWithSub
        to="/crafted/widgets"
        title="Widgets"
        icon={gen022}
      >
        <AsideMenuItem to="/crafted/widgets/lists" title="Lists" hasBullet={true} />
        <AsideMenuItem to="/crafted/widgets/statistics" title="Statistics" hasBullet={true} />
        <AsideMenuItem to="/crafted/widgets/charts" title="Charts" hasBullet={true} />
        <AsideMenuItem to="/crafted/widgets/mixed" title="Mixed" hasBullet={true} />
        <AsideMenuItem to="/crafted/widgets/tables" title="Tables" hasBullet={true} />
        <AsideMenuItem to="/crafted/widgets/feeds" title="Feeds" hasBullet={true} />
      </AsideMenuItemWithSub>
      <div className="menu-item">
        <div className="menu-content pt-8 pb-2">
          <span className="menu-section text-muted text-uppercase fs-8 ls-1">Apps</span>
        </div>
      </div>
      <AsideMenuItemWithSub
        to="/apps/chat"
        title="Chat"
        icon={com012}
      >
        <AsideMenuItem to="/apps/chat/private-chat" title="Private Chat" hasBullet={true} />
        <AsideMenuItem to="/apps/chat/group-chat" title="Group Chart" hasBullet={true} />
        <AsideMenuItem to="/apps/chat/drawer-chat" title="Drawer Chart" hasBullet={true} />
      </AsideMenuItemWithSub>
      <AsideMenuItem
        to="/apps/user-management/users"
        icon={gen051}
        title="User management"
      />
      <div className="menu-item">
        <div className="menu-content">
          <div className="separator mx-1 my-4"></div>
        </div>
      </div>
      <div className="menu-item">
        <a
          target="_blank"
          className="menu-link"
          href={process.env.REACT_APP_PREVIEW_DOCS_URL + "/docs/changelog"}
        >
          <span className="menu-icon">
            <Image src={gen005} className="svg-icon-2" alt={"gen005"} />
          </span>
          <span className="menu-title">Changelog {process.env.NEXT_PUBLIC_APP_VERSION}</span>
        </a>
      </div>
    </>
  );
}
