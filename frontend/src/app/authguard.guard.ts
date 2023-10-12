import { Injectable } from '@angular/core';
import { CanActivate, Router } from '@angular/router';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard{
  constructor(private router: Router) {}

  canActivate(): boolean {
    // Check if the user is authenticated (e.g., by checking the presence of the JWT token)
    const token = localStorage.getItem('token');
    if ((token) && token != undefined) {
      return true;
    } else {
      // If not authenticated, redirect to the login page
      this.router.navigate(['/login']); // Replace '/login' with your actual login path
      return false;
    }
  }
}