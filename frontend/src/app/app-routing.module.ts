import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { HomeComponent } from './components/home/home.component';
import { AuthGuard } from '../app/authguard.guard'; 
import { LayoutComponent } from './layout/layout.component';
import { PredictionsComponent } from './components/predictions/predictions.component';
import { OriginaldataComponent } from './components/originaldata/originaldata.component';

const routes: Routes = [
  { path: '', redirectTo: '/home', pathMatch: 'full' },
  {
  path: '',
  component: LayoutComponent,
  children: [
    { path: 'home', component: HomeComponent, canActivate: [AuthGuard] },
     { path: 'predictions', component: PredictionsComponent, canActivate: [AuthGuard] },
     { path: 'original_data', component: OriginaldataComponent, canActivate: [AuthGuard] },
  ]},
  {path: 'login', component: LoginComponent},
  {path: 'register', component: RegisterComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
