const cart = document.querySelector('.shopping-cart');

document.querySelector('#cart').onclick = () => {
  cart.classList.toggle('active');
  login.classList.remove('active');
}

const login = document.querySelector('.login-form');

document.querySelector('#login').onclick = () => {
  login.classList.toggle('active');
  cart.classList.remove('active');
}

const navbar = document.querySelector('.navbar');

document.querySelector('#menu').onclick = () => {
  navbar.classList.toggle('active');
  cart.classList.remove('active');
  login.classList.remove('active');
}

window.onscroll = () => {
  cart.classList.remove('active');
  login.classList.remove('active');
  navbar.classList.remove('active');
}

const previewContainer = document.querySelector('.menu-preview-container');
const previewBox = previewContainer.querySelectorAll('.menu-preview');

document.querySelectorAll('.menu .box').forEach(menu => {
  menu.onclick = () => {
    previewContainer.style.display = 'flex';
    const name = menu.getAttribute('data-name');
    previewBox.forEach(preview => {
      const target = preview.getAttribute('data-target');
      if(name == target){
        preview.classList.add('active');
      }
    });
  };
});

previewContainer.querySelector('#close').onclick = () => {
  previewContainer.style.display = 'none';
  previewBox.forEach(close => {
    close.classList.remove('active');
  });
};
