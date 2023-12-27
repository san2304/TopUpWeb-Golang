document.addEventListener('DOMContentLoaded', function () {
    // active nav link
    var ul = document.querySelector('nav ul');
    var li = document.querySelectorAll('nav ul li');

    li.forEach(el => {
        el.addEventListener('click', function () {
            ul.querySelector('.active').classList.remove('active');
            el.classList.add('active');
        });
    });
    // image slider
    const swiper = new Swiper('.swiper-container', {
        autoplay: {
            delay: 3000,
            disableOnInteraction: false,
        },
        loop: true,
        pagination: {
            el: '.swiper-pagination',
            clickable: true,
        },
        preloadImages: false,
        lazy: true,
    });
});