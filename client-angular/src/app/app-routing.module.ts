import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './auth/login/login.component';
import { RegisterComponent } from './auth/register/register.component';
import { ChatComponent } from './chat/chat.component';
import { FriendsComponent } from './friends/friends.component';
import { AuthOptionalGuard } from './guard/auth-optional.guard';
import { AuthGuard } from './guard/auth.guard';
import { MessagesComponent } from './messages/messages.component';
import { ChangePasswordComponent } from './profile/change-password/change-password.component';
import { UpdateProfileComponent } from './profile/update-profile/update-profile.component';
import { UserPageComponent } from './profile/user-page/user-page.component';

const routes: Routes = [
    {
      path: 'user',
      component: UserPageComponent,
      pathMatch: 'full',
      canActivate: [AuthGuard],
    },
    {
        path: 'user/:id',
        component: UserPageComponent,
        canActivate: [AuthOptionalGuard]
      },
    {
      path: 'messages',
      component: MessagesComponent,
      canActivate: [AuthGuard]
    },
    {
      path: 'chat',
      component: ChatComponent,
      canActivate: [AuthGuard]
    },
    {
      path: 'friends',
      component: FriendsComponent,
      canActivate: [AuthGuard]
    },
    {
      path: 'friends/:id',
      component: FriendsComponent,
    },
    {
      path: 'edit-profile',
      component: UpdateProfileComponent,
      canActivate: [AuthGuard]
    },
    {
      path: 'login',
      component: LoginComponent
    },
    {
      path: 'register',
      component: RegisterComponent
    },
    {
      path: 'change-password',
      component: ChangePasswordComponent,
      canActivate: [AuthGuard]
    },
    {
      path: '',
      redirectTo: 'login',
      pathMatch: 'full'
    },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }