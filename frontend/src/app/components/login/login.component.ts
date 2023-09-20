import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  email: string = '';
  password: string = '';

  constructor(private router: Router) {}

  onSubmit() {
    // Perform your login logic here (e.g., send login request to the server)

    // After successful login, you can navigate to the home page or any other route
    // For example, navigate to '/home':
    this.router.navigate(['/home']);
  }
}