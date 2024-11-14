import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export function middleware(request: NextRequest) {
  const authUrls: string[] = ["/login"];
  const tld = request.nextUrl.protocol + "//" + request.nextUrl.hostname + ":" + request.nextUrl.port;
  const currentPathname = request.nextUrl.pathname;
  const hasAccessToken = request.cookies.has("AccessToken");

  if (!authUrls.includes(currentPathname) && !hasAccessToken) {
    return NextResponse.redirect(new URL("/login", request.url));
  }

  if (authUrls.includes(currentPathname) && hasAccessToken) {
    return NextResponse.redirect(new URL("/dashboard", tld));
  }

  return NextResponse.next();
}

export const config = {
  matcher: "/((?!api|static|.*\\..*|_next).*)",
};
