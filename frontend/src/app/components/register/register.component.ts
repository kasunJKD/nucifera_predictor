import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { catchError } from 'rxjs/operators';
import { Observable, throwError } from 'rxjs';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  registrationForm: FormGroup;

  constructor(
    private fb: FormBuilder,
    private http: HttpClient // Inject HttpClient
  ) {
    this.registrationForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      confirmPassword: ['', [Validators.required]]
    });
  }

  onSubmit() {
    if (this.registrationForm.valid) {
      const formData = this.registrationForm.value;

      // Check if passwords match
      if (formData.password !== formData.confirmPassword) {
        // Handle password mismatch (e.g., display an error message)
        console.error('Passwords do not match');
        return;
      }

      // Make the POST request to the backend
      this.http.post('http://localhost:11001/api/membership/signUp', formData)
        .pipe(
          catchError(this.handleError)
        )
        .subscribe(
          (response) => {
            // Handle success response here
            console.log('Registration successful:', response);
            // You can also navigate to a login page or show a success message to the user
          }
        );
    }
  }

  private handleError(error: any): Observable<any> {
    // Handle error response here
    console.error('Registration error:', error);
    // You can display an error message to the user
    return throwError(error);
  }
}


