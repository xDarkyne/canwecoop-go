export interface User {
  ID: string;
  DisplayName: string;
  AvatarUrl: string;
  ProfileUrl: string;
  IsTester: boolean;
  IsAdmin: boolean;
  Friends: [];
  CreatedAt: string;
  LastLoggedIn: string;
}
